# Integers

Go has `int`, `int8/16/32/64`, `uint`, `uint8/16/32/64`, `uintptr`.
Unlike Rust there is no `u128`, and `int` is 32 or 64 bits depending on platform.
Go never implicitly converts between integer types: `uint32(m)` is required.

## Task

Use `multiplier` (a `uint8`) in the computation. Convert it to `uint32` first.

## Check

```bash
go test ./exercises/02_calculator/01_integers/ -v
```
