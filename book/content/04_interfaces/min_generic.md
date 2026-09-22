---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Generic Min"
weight: 7
draft: false
---

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

---

*Source: `exercises/04_interfaces/07_min_generic/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
