# Conversions

The integers lesson laid down the law: no implicit conversion, ever. This
exercise is the other half of that law — the explicit machinery for moving
values between types, and the two very different fates awaiting them.

## The one syntax

Every numeric conversion in Go looks identical:

```go
// Syntax: <Type>(<value>)
int64(x) // x, whatever its old type, viewed as int64
```

No keywords, no cast operators, no "convert" builtin — the type name *is*
the operation. Once you've seen `uint32(multiplier)`, you've seen them all.
Uniformity like this is deliberate: conversions should be visually boring,
because boring code gets reviewed correctly.

## Widening: always safe

Moving to a roomier type preserves the value, unconditionally:

```go
func ToInt64(x int32) int64 {
    return int64(x)
}
```

Every `int32` fits in an `int64`, so `ToInt64(-42)` is `-42` by
construction. Widen-then-compute is the standard prelude to any arithmetic
that might overflow the narrow type — the overflow lesson's wide-type
strategy in miniature.

## Narrowing: bits fall off

The reverse direction keeps only what fits:

```go
uint8(300)   // 44  — low 8 bits survive
int8(200)    // -56 — reinterpreted as signed
int(2.9)     // 2   — fraction discarded, not rounded
```

None of these fail. Go assumes you measured before you cut — and if you
didn't, you get the overflow lesson's wrap wearing a conversion's clothes.
Before every narrowing conversion, ask: *can this value exceed the target?*
If yes, check first (the `sum < a` family) or clamp (the saturating
family). The detector always transfers; only the types change.

One conversion deserves a warning sign of its own: `string(65)` is `"A"`,
not `"65"` — it converts to the Unicode character with that code point.
Turning numbers into digits is the `strconv` package's job, arriving with
the strings section. If a conversion ever surprises you, suspect the
*meaning* changed, not just the width.

## Task

Fix `ToInt64` in `calc.go` to convert `x` with `int64(x)`.

## Check

```bash
go test ./exercises/02_calculator/10_conversions/ -v
```
