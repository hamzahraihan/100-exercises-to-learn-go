# Overflow

Your calculator has been suspiciously well-behaved: every sum fit, every test
told the truth. Time to break that streak — on purpose.

The exercise in `exercises/02_calculator/08_overflow` asks an uncomfortable
question: what happens when the answer doesn't fit the type? The full
integer model could fill a chapter on its own. You'll get the slice that
matters — spotting the wrap and picking a fix — and leave the rest for the
day you meet it in the wild.

## A number that doesn't fit

A `uint8` holds 8 bits, so it represents the values 0 to 255. Now add:

```go
var a uint8 = 200
var b uint8 = 100
sum := a + b // 44?!
```

200 + 100 is 300, and 300 is not a `uint8`. Go doesn't stop your program
over this. It wraps around modulo 256 and carries on: `(200 + 100) % 256`
is 44. No error, no warning — just a wrong number wearing a correct type.

## The broken contract

This is what **integer overflow** means: the mathematically correct result
of an operation doesn't fit the type it's supposed to come back in. Addition
promises "two `uint8` values in, one `uint8` out" — and 300 breaks the deal.

Overflows bite because everything *looks* fine. Lengths, counters, and
indexes computed this way go wrong quietly, far from the line that caused
it. That's why every narrow unsigned addition deserves a second thought.

> Subtraction wraps too: `uint8(0) - 1` is 255, not -1. That event is
> called **underflow**. Everything below applies to it as well, but we'll
> talk only about addition — it's what this exercise tests.

## Why not just use a bigger type?

One tempting fix: let the language promote the result automatically. `a + b`
overflows `uint8`, so why not quietly hand back a `uint16`?

Because Go refuses to guess about width. The type of `a + b` where both are
`uint8` is `uint8` — always, with no exceptions. Widening is your explicit
decision to make, visible in the code:

```go
w := uint16(a) + uint16(b) // 300, honestly held
```

No hidden costs, no surprise types flowing downstream. If you want room,
say so up front.

## Two ways out

Since promotion won't save us, what can we do when a sum might overflow?
It boils down to two approaches:

- **Detect it** — compute in the narrow type, notice the wrap, and tell
  the caller via a flag.
- **Sidestep it** — compute in a wider type and check the range yourself.

### Detect it

Without wrapping, `a + b` is at least as big as `a`. With wrapping, 256 is
subtracted behind your back, so the result comes out *smaller* than `a`.
That single comparison is the whole detector:

```go
// Syntax: compute, then compare against an input
func AddUint8(a, b uint8) (sum uint8, overflow bool) {
    sum = a + b
    overflow = sum < a // wrapped if and only if the result shrank
    return sum, overflow
}
```

Run the test cases through it by hand:

```go
AddUint8(100, 100) // (200, false) — 200 fits, 200 < 100 is false
AddUint8(200, 100) // (44, true)   — wrapped, 44 < 200 is true
AddUint8(255, 1)   // (0, true)    — the classic edge, 0 < 255 is true
```

Returning the value *and* a flag — `(sum, overflow)` — is the idiomatic Go
shape for "here's the answer, and here's whether to trust it." You'll meet
its cousins soon: map lookups return `(value, ok)`, type assertions return
`(value, ok)`. Same rhythm every time.

### Sidestep it

When you need the true total rather than a yes/no flag, widen first and
compare against the boundary:

```go
func AddUint8Wide(a, b uint8) (uint8, bool) {
    w := uint16(a) + uint16(b)
    if w > 255 {
        return uint8(w), true
    }
    return uint8(w), false
}
```

Order matters: convert *before* adding. `uint8(x + y)` computed from two
`uint8` values wraps before the conversion ever runs — the damage is done
upstream of your check.

For machine-word arithmetic the standard library helps: `math/bits` offers
`Add` for `uint` plus `Add32` and `Add64` sized variants, each returning a
sum and a carry flag. There is no 8-bit variant, which is exactly why the
one-line comparison above is the usual choice for `uint8`.

Which approach should you reach for? If the caller can proceed with a flag
— retry, clamp, report — detect. If the caller needs the real number,
widen. Either way, never return a quiet wrong answer when a loud signal
was one comparison away.

## Task

Fix `AddUint8` in `calc.go` to report wraparound via the `overflow` flag.

## Check

```bash
go test ./exercises/02_calculator/08_overflow/ -v
```
