# Generic Min

Go has no function overloading. No `Min` for ints beside another `Min` for
strings — one package, one name, one signature. Every language without
overloads faces the same fork: duplicate the function per type, or abstract
over the difference. This exercise takes the second road.

## One name, every ordered type

```go
func Min[T cmp.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}
```

`cmp.Ordered` — from the standard `cmp` package — permits every type that
supports `<`: all the ints and floats, plus strings. The constraint is what
makes the body legal: `<` compiles because *no permitted type lacks it*.
Same permission-slip logic as the generic sum, one section earlier; new
this time is that the callers span kinds, not just widths:

```go
Min(3, 7)   // 3 — T is int, inferred
Min(7, 3)   // 3 — argument order never matters to the answer
Min("b", "a") // "a" — strings compare lexicographically, byte by byte
```

Inference again fills `T` silently — `Min(3, 7)` never spells `int`, and
`Min("b", "a")` never spells `string`. One function, every ordered type,
call sites indistinguishable from ordinary calls.

## What "ordered" buys and doesn't

`<` on numbers is numeric; on strings it's lexicographic byte order
(`"b" > "a"`, `"Z" < "a"` — capitals sort before lowercase in UTF-8, a fact
that surprises exactly once). The constraint promises comparability, not
intuition — string ordering follows bytes, and bytes follow history.

And the boundary holds in both directions: try `Min` on a struct or a slice
and the call fails, because neither is `cmp.Ordered`. Constraints cut both
ways by design — everything permitted works, everything else is rejected up
front. When a future function needs equality instead of order, `comparable`
from the previous lesson waits. Pick the weakest constraint the body needs;
each one is a precise statement about what the code actually does.

## Task

Complete `Min` in `min.go` so it returns the smaller of `a` and `b`
for any ordered type.

## Check

```bash
go test ./exercises/04_interfaces/07_min_generic/ -v
```
