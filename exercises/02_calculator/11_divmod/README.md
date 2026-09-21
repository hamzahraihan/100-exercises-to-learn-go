# Div Mod

Go functions can return multiple values, so quotient and remainder come back
together: `func DivMod(a, b int) (int, int)`. Integer `/` truncates toward
zero and `%` gives the remainder with the sign of the dividend.

## Task

Fix `DivMod` in `calc.go` to return `a / b` and `a % b`.

## Check

```bash
go test ./exercises/02_calculator/11_divmod/ -v
```
