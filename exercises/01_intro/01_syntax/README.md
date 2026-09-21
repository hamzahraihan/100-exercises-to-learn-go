# Syntax

Go functions look like this:

```go
func Compute(a, b int) int {
    return a + b
}
```

`package syntax` declares the package. `func` declares a function,
`(a, b int)` are typed parameters, the trailing `int` is the return type.

## Task

Fix `Compute` in `syntax.go` so the test passes.

## Check

```bash
go test ./exercises/01_intro/01_syntax/ -v
```
