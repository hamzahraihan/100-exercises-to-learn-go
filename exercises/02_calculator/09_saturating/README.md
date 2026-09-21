# Saturating Add

Saturating arithmetic clamps at the type's maximum instead of wrapping:
`SaturatingAdd(200, 100)` on `uint8` gives 255, not 44. Since Go wraps
silently, you detect the wrap (`sum < a`) and substitute the max value.

## Task

Fix `SaturatingAdd` in `calc.go` to return 255 when the sum overflows.

## Check

```bash
go test ./exercises/02_calculator/09_saturating/ -v
```
