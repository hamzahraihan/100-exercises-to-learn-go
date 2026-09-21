# Any

`any` is an alias for `interface{}`: the empty interface that every type
implements. A parameter of type `any` accepts a value of any type, which
makes it handy for generic formatting, logging, or test helpers:

```go
func SprintAny(v any) string {
    // ... format v with the fmt package ...
}
```

Use formatting verbs such as `%v` (or helpers like `fmt.Sprint`) to turn
an unknown value into a string. Reach for `any` when you truly accept
anything; prefer a concrete type or a narrower interface when you know
more about what callers will pass, since `any` gives up compile-time
checking.

## Task

Implement `SprintAny` in `any.go` so it returns the formatted value and
the test passes.

## Check

```bash
go test ./exercises/04_interfaces/03_any/ -v
```
