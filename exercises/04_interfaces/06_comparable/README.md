# Comparable

Maps are Go's dictionaries — and every dictionary needs a rule for telling
keys apart. That rule is `==`, and not every type qualifies: slices and
maps can't be compared with `==` at all. This exercise writes one function
over all legal key types, with the compiler enforcing the boundary.

## Keys must compare

```go
// Syntax: func <Name>[K comparable, V any](m map[K]V) ...
func Keys[K comparable, V any](m map[K]V) []K {
    keys := make([]K, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}
```

Two parameters, two different jobs. `K` carries the **`comparable`**
constraint — the predeclared bound matching exactly the types that support
`==` and `!=`. Without it, `K` could be a slice type and the map itself
would be illegal; the constraint makes "valid map key" a compile-time fact
rather than documentation. `V` stays `any` because values are never
compared here — why demand more than the body uses?

`make([]K, 0, len(m))` pre-sizes the result: length 0, capacity for every
key, one allocation instead of many. `append` grows into it. And `for k :=
range m` walks keys only — the single-variable range form, sibling to the
index-skipping `for _, v := range vals` from the calculator section.

## Why the test looks like that

```go
Keys(map[string]int{"a": 1}) // single key
Keys(map[int]string{})       // empty map
```

One key, then none. Conspicuously absent: a three-key map asserting order.
That's deliberate — **map iteration order is randomized** in Go, on purpose,
so programs can't accidentally depend on it. A multi-key assertion would be
flaky by design. Single-element and empty cases test the logic; order is
left untested because it's unpromised.

Same reason for `reflect.DeepEqual` on the result: slices don't support
`==` (they're not `comparable` — the lesson, applied to the test itself),
so deep comparison does the checking. The constraint you just learned
explains the test you just read.

## Task

Implement `Keys` in `keys.go` so it returns every key in `m` and the
test passes.

## Check

```bash
go test ./exercises/04_interfaces/06_comparable/ -v
```
