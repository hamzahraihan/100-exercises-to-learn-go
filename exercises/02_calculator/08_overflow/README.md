# Overflow

Don't jump ahead!
Make sure the earlier calculator exercises pass before starting this one.
It's located in `exercises/02_calculator/08_overflow`.

So far every addition you've written has fit comfortably in its type. This
exercise looks at what happens when it doesn't. We won't cover the whole
integer model — *just enough* to detect a wrap and choose a fix. One step
at a time!

## Integer types

Go gives you sized unsigned integers: `uint8`, `uint16`, `uint32`, `uint64`
(and the matching signed `int8` … `int64`). The number is the bit width.

A `uint8` holds 8 bits, so it represents the values 0 to 255. That range is
the whole point of the type — and the source of this exercise:

```go
var max uint8 = 255
```

## What happens on overflow

Adding past the top does **not** stop your program. Go wraps around modulo
2ⁿ (256, for `uint8`) and keeps going silently:

```go
var a uint8 = 200
var b uint8 = 100
sum := a + b // 44, not 300: (200 + 100) % 256 == 44
```

The result is well-defined but almost never what you wanted. Counters,
lengths, and indexes computed this way go wrong quietly, which is why every
small unsigned addition deserves a second thought.

## Detecting overflow

For addition, the wrap test is a single comparison. Without wrapping,
`a + b` is at least `a`. With wrapping, 256 is subtracted, so the result
is smaller than `a`:

```go
func AddUint8(a, b uint8) (sum uint8, overflow bool) {
    sum = a + b
    overflow = sum < a // wrapped if and only if the result shrank
    return sum, overflow
}
```

Walk through the cases from the test suite:

```go
AddUint8(100, 100) // (200, false) — 200 fits in a uint8
AddUint8(200, 100) // (44, true)   — wrapped: 44 < 200
AddUint8(255, 1)   // (0, true)    — the classic edge: 0 < 255
```

It is considered idiomatic to report the situation rather than hide it:
return the wrapped value *and* a flag, `(sum, overflow)`, and let the caller
decide what to do. You will meet this two-value shape again with maps
(`value, ok`) and type assertions.

## Choosing a strategy

Checking after the fact is one option. Avoiding the narrow type is another —
compute in a wider type, then test the range:

```go
func AddUint8Wide(a, b uint8) (uint8, bool) {
    w := uint16(a) + uint16(b)
    if w > 255 {
        return uint8(w), true
    }
    return uint8(w), false
}
```

The key detail is converting *before* adding. Converting after the addition
is too late — the wrap has already happened.

For machine-word arithmetic the standard library helps: `math/bits` offers
`Add` (for `uint`) with `Add32` and `Add64` sized variants, each returning
the sum and a carry flag. There is no 8-bit variant, which is why the
explicit comparison above is the usual choice for `uint8`.

## Common mistakes

- **Converting too late:** `uint8(x + y)` where `x` and `y` are already
  `uint8` wraps before the conversion runs. Widen the inputs first.
- **Reusing the unsigned test for signed types:** `sum < a` only works for
  unsigned addition. `int8(120) + 20` wraps to a negative number, which
  needs a range check *before* adding instead.
- **Ignoring the flag:** computing `overflow` and then not branching on it
  leaves the bug in place. If the caller can't act on it, prefer failing
  loudly (return an error) over returning a quiet wrong number.

## Task

Fix `AddUint8` in `calc.go` to report wraparound via the `overflow` flag.

## Check

```bash
go test ./exercises/02_calculator/08_overflow/ -v
```
