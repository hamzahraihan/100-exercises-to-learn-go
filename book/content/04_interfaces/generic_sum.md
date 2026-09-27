---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Generic Sum"
weight: 2
draft: false
---

# Generic Sum

Summing `[]int64` and summing `[]float64` is the same loop with different
labels. Without generics you'd write it twice — `SumInt64`, `SumFloat64` —
and every fix would need applying twice, drifting apart by the third copy.
This exercise writes it once, for every type on an approved list.

## One function, many types

```go
// Syntax: func <Name>[<param> <constraint>](<args>) <results>
func Sum[T int64 | float64](vals []T) T {
    var total T
    for _, v := range vals {
        total += v
    }
    return total
}
```

`T` is a **type parameter**: a placeholder the caller fills in. The
**constraint** after it — `int64 | float64` — is the approved list. `Sum`
accepts `[]int64` or `[]float64` and nothing else; try `[]string` and the
*call* fails to compile, before a single test runs.

The constraint is also a permission slip for the body. `total += v` compiles
because *every* permitted type supports `+=`. If the list admitted a type
without addition, the body would break for all callers — so the compiler
checks the body against the constraint, not against each future caller.
Write the loop once, prove it against the list, trust it everywhere.

## Calling without ceremony

Nobody writes the type argument here:

```go
Sum([]int64{1, 2, 3})   // T inferred as int64 → 6
Sum([]float64{1.5, 2.5}) // T inferred as float64 → 4.0
```

**Inference** fills `T` from the argument — the compiler sees `[]int64`
and instantiates that version silently. Explicit arguments
(`Sum[int64](...)`) exist for the ambiguous cases, but idiomatic calls
leave them out. Same function, two compiled variants, zero duplication.

And `var total T`? The zero-value promise, generalized. Whatever `T` turns
out to be, `total` starts as its zero — `0` here, `0.0` there — ready to
accumulate. The variables lesson's rule never changed; it just learned a
new letter.

## Task

Complete `Sum` in `sum.go` by ranging over `vals` and accumulating into
`total` so the test passes for both `int64` and `float64` inputs.

## Check

```bash
go test ./exercises/04_interfaces/02_generic_sum/ -v
```

---

*Source: `exercises/04_interfaces/02_generic_sum/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
