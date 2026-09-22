---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Generic Sum"
weight: 2
draft: false
---

# Generic Sum

Go generics let one function work for several types. A type parameter with
a constraint lists the permitted types:

```go
func Sum[T int64 | float64](vals []T) T {
    var total T
    // ... range over vals and accumulate into total ...
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

---

*Source: `exercises/04_interfaces/02_generic_sum/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
