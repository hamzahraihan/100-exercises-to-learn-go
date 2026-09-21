# Sum To

Go has only `for`: there is no `while` keyword. The same C-style header
`for i := 1; i <= n; i++` covers counting loops, and `for condition { }`
covers while-style loops.

## Task

Fix `SumTo` in `calc.go` to return `1 + 2 + ... + n` (0 for `n <= 0`).

## Check

```bash
go test ./exercises/02_calculator/06_sum_to/ -v
```
