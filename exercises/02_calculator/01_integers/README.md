# Integers

Go doesn't have *an* integer type. It has a whole menu — and the compiler
makes you order precisely. This exercise's stub looks innocent: multiply and
add. It doesn't compile as written, and the reason is the entire lesson.

## The menu

```go
int   uint     // platform-sized: 32 bits on 32-bit systems, 64 on 64-bit
int8  int16  int32  int64
uint8 uint16 uint32 uint64
```

Signed or unsigned, then a width. `uint8` holds 0–255 (you've met its
wrapping temper in the overflow exercise); `int64` stretches past nine
quintillion in both directions. `int` and `uint` follow the platform —
write `int` for everyday counting and only reach for a width when the
protocol, the file format, or memory demands it.

## Types don't mix

Here's the stub's trap:

```go
var multiplier uint8 = 4
func Compute(a, b uint32) uint32 {
    return a + b  // TODO: picks up multiplier... how?
}
```

The natural try — `a + b*multiplier` — is rejected. `b` is `uint32`,
`multiplier` is `uint8`, and Go performs **no implicit conversion**, ever.
No promotion, no quiet widening. Operands of `*` must already share a type,
or the program doesn't build.

This strictness is the same instinct behind the overflow lesson: the
language never silently changes your numbers' width. If a value moves
between types, the code says so out loud.

## Conversions say it out loud

Moving a value between integer types is a **conversion**, written as the
target type around the value:

```go
// Syntax: <Type>(<value>)
uint32(multiplier) // uint8(4) becomes uint32(4)
```

Conversions go both directions. Widening (`uint8` → `uint32`) always
preserves the value. Narrowing (`uint32` → `uint8`) keeps only the low bits
— `uint8(300)` is 44, the same wrap a later exercise will teach you to
detect. The syntax never warns you; knowing the direction is your job.

So the fix reads almost like a sentence: convert the multiplier to the
working type, *then* do arithmetic:

```go
return a + b*uint32(multiplier)
```

## Task

Use `multiplier` (a `uint8`) in the computation. Convert it to `uint32` first.

## Check

```bash
go test ./exercises/02_calculator/01_integers/ -v
```
