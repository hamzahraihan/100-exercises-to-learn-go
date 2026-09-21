# Generic Min

One generic function replaces a family of overloads. The `cmp.Ordered`
constraint (from the `cmp` package) permits any type that supports `<`:
numbers and strings:

```go
func Min[T cmp.Ordered](a, b T) T {
    // ... compare a and b and return the smaller ...
}
```

Because the constraint guarantees ordering, the body can compare `a`
and `b` directly and works unchanged for `int`, `float64`, `string`,
and every other ordered type.

## Task

Complete `Min` in `min.go` so it returns the smaller of `a` and `b`
for any ordered type.

## Check

```bash
go test ./exercises/04_interfaces/07_min_generic/ -v
```
