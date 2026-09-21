# If Else

Go's `if/else` looks familiar but drops the parentheses around the condition,
and the braces are required even for one-line bodies: `if a >= b { ... } else { ... }`.

## Task

Fix `Max` in `calc.go` to return the larger of `a` and `b` with an `if/else`.

## Check

```bash
go test ./exercises/02_calculator/03_if_else/ -v
```
