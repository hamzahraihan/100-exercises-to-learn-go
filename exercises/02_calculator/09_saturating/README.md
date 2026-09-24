# Saturating Add

The overflow exercise taught you to *report* the wrap. But sometimes the
caller can't do anything with a flag — a volume knob, a progress bar, a
pixel brightness has nowhere to put "error." For those, there's a third
answer between wrapping and panicking: **clamp**.

## Clamping instead of wrapping

**Saturating** arithmetic pins the result at the type's boundary instead of
wrapping around it:

```go
SaturatingAdd(10, 20)  // 30  — fits, plain addition
SaturatingAdd(200, 100) // 255 — would wrap to 44, clamped at MaxUint8
```

200 + 100 is 300; the honest `uint8` answer doesn't exist, so 255 — the
nearest representable value — stands in. Audio mixers, image filters, and
game health bars all saturate: better slightly wrong in the right direction
than catastrophically wrapped.

## Reusing the detector

You already own the hard part. The overflow check from two exercises ago
detects exactly the moment to clamp:

```go
func SaturatingAdd(a, b uint8) uint8 {
    sum := a + b
    if sum < a {
        return math.MaxUint8 // 255: the ceiling, from package math
    }
    return sum
}
```

`math.MaxUint8` names the boundary instead of burying a magic 255 in the
code. Named limits document intent (`MaxUint8` says *ceiling of the type*;
`255` says *a number someone remembered*) and survive intact if the type
ever widens. The standard library provides the full set — `MaxInt`,
`MaxInt64`, `MaxFloat64`, and friends — so reach for them before typing a
literal boundary.

Notice the structure: detect, then substitute. Saturating arithmetic is
*overflow handling with the handling filled in* — same diagnosis, different
prescription. When the next exercise hands you a new policy for the same
wrap, the detector won't change; only the `if` body will.

## Task

Fix `SaturatingAdd` in `calc.go` to return 255 when the sum overflows.

## Check

```bash
go test ./exercises/02_calculator/09_saturating/ -v
```
