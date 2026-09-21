# Panics

Go signals unrecoverable errors with `panic`. Calling `panic("message")` stops
normal execution and unwinds the stack; tests recover from it with `recover()`.
Unlike returning an error, a panic crashes the program unless recovered.

## Task

Fix `Divide` in `calc.go` to `panic("division by zero")` when `b == 0`,
otherwise return `a / b`.

## Check

```bash
go test ./exercises/02_calculator/04_panics/ -v
```
