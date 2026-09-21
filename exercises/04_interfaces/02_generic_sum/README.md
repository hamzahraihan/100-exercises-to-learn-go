# Generic Sum

Go generics let one function work for several types. A type parameter with
a constraint lists the permitted types:

```go
func Sum[T int64 | float64](vals []T) T {
    var total T
    for _, v := range vals {
        total += v
    }
    return total
}
```

The constraint `int64 | float64` plays the role that trait bounds play in
the Rust course this section is adapted from: it says which types callers
may plug in while still allowing `+=` on values of type `T`.

## Task

Complete `Sum` in `sum.go` by ranging over `vals` and accumulating into
`total` so the test passes for both `int64` and `float64` inputs.

## Check

```bash
go test ./exercises/04_interfaces/02_generic_sum/ -v
```
