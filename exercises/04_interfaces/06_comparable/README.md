# Comparable

Map keys must support `==` and `!=`. The predeclared `comparable`
constraint captures exactly that set of types, so a generic function
over map keys uses it for `K`:

```go
func Keys[K comparable, V any](m map[K]V) []K {
    // ... range over m and collect each key ...
}
```

`V` stays `any` because values never need comparing here. Without the
`comparable` bound the compiler could not guarantee `K` is a legal map
key.

## Task

Implement `Keys` in `keys.go` so it returns every key in `m` and the
test passes.

## Check

```bash
go test ./exercises/04_interfaces/06_comparable/ -v
```
