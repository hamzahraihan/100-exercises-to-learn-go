# Floats

`float64` holds fractional numbers. Converting after integer division keeps
the truncation: `float64(5/2)` is `2.0`, while `float64(5)/2` is `2.5`.
The `math` package (`math.Pi`, `math.Sqrt`) covers the rest.

## Task

Fix `Half` in `calc.go` so odd inputs keep their `.5`.

## Check

```bash
go test ./exercises/02_calculator/12_floats/ -v
```
