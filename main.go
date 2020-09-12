package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/stateful"
	"github.com/alecthomas/repr"
)

var (
	Lexer = lexer.Must(stateful.New(stateful.Rules{
		"Root": {
			{"Ident", `\w+`, nil},
			{"String", `"`, stateful.Push("String")},
		},
		"String": {
			{"StringEnd", `"`, stateful.Pop()},
			{"Escaped", `\\.`, nil},
			{"Interpolated", `\${`, stateful.Push("Interpolated")},
			{"Char", `[^"\\]+`, nil},
		},
		"Interpolated": {
			{"End", `}`, stateful.Pop()},
			stateful.Include("Root"),
		},
	}))

	Parser = participle.MustBuild(
		&AST{},
		participle.Lexer(Lexer),
	)
)

type AST struct {
	Pos  lexer.Position
	Expr *Expr `parser:"@@"`
}

type Expr struct {
	Pos    lexer.Position
	Ident  *Ident  `parser:"( @@"`
	String *String `parser:"| @@ )"`
}

type Ident struct {
	Pos  lexer.Position
	Text string `parser:"@Ident"`
}

type String struct {
	Pos       lexer.Position
	Fragments []*StringFragment `parser:"String @@ StringEnd"`
}

type StringFragment struct {
	Pos          lexer.Position
	Escaped      *string       `parser:"( @Escaped"`
	Interpolated *Interpolated `parser:"| @@"`
	Text         *string       `parser:"| @Char )"`
}

type Interpolated struct {
	Pos  lexer.Position
	Expr *Expr `parser:"Interpolated @@? End"`
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %s", err)
		os.Exit(1)
	}
}

func run() error {
	ast := &AST{}
	err := Parser.ParseString(`"echo $HOME ${foo}"`, ast)
	repr.Println(ast)
	return err
}
