# Conversions

Go never converts numbers implicitly: assigning an `int32` to an `int64`
needs an explicit conversion written as `T(x)`, e.g. `int64(x)`. There is no
`as` cast keyword. Widening conversions like `int32` to `int64` always
preserve the value.

## Task

Fix `ToInt64` in `calc.go` to convert `x` with `int64(x)`.

## Check

```bash
go test ./exercises/02_calculator/10_conversions/ -v
```
