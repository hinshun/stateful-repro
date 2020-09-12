# stateful-repro

```sh
❯ go run .
&main.AST{
  Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
  Expr: &main.Expr{
    Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
    String: &main.String{
      Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
      Fragments: []*main.StringFragment{
        &main.StringFragment{
          Pos: Position{Filename: "", Offset: 1, Line: 1, Column: 2},
          Text: &"echo $HOME ${foo}",
        },
      },
    },
  },
}
```
