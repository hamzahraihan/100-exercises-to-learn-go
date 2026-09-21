# Overflow

Unsigned arithmetic in Go wraps silently on overflow: `uint8(200) + 100`
gives 44 with no panic (unlike Rust, which panics in debug builds).
Idiomatic Go checks explicitly: if `sum < a`, the addition wrapped around.

## Task

Fix `AddUint8` in `calc.go` to report wraparound via the `overflow` flag.

## Check

```bash
go test ./exercises/02_calculator/08_overflow/ -v
```
