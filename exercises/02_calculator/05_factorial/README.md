# Factorial

Go has a single loop keyword, `for`, which handles every loop shape.
A counting loop looks like `for i := 1; i <= n; i++ { ... }`.
Start the accumulator at 1 so that `Factorial(0)` is 1.

## Task

Fix `Factorial` in `calc.go` to return `n!` using a `for` loop.

## Check

```bash
go test ./exercises/02_calculator/05_factorial/ -v
```
