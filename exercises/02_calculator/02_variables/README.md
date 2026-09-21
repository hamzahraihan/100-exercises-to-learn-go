# Variables

Go offers two ways to declare a variable: `var x int = 1` states the type
explicitly, while `x := 1` lets the compiler infer it. The short form `:=`
only works inside functions.

## Task

Fix `Double` in `calc.go` to return twice `x`, using `:=` to declare the result.

## Check

```bash
go test ./exercises/02_calculator/02_variables/ -v
```
