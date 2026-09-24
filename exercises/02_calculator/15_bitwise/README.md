# Bitwise operators

Fourteen exercises of arithmetic, and numbers have behaved like quantities.
Now they become *switchboards*: individual bits flipped, combined, and
tested. The calculator section closes with the lowest-level arithmetic Go
offers — and the strangest-looking constants you've met.

## Bits as flags

A `uint8` holds eight independent on/off switches. The operators manipulate
them directly:

```go
a | b  // OR:  a switch is on if it's on in either side
a & b  // AND: a switch is on only if it's on in both sides
a ^ b  // XOR: a switch is on if the sides disagree
a &^ b // AND NOT: switches on in a but off in b (Go's "clear" operator)
a << n // shift left n places (multiplies by 2ⁿ)
a >> n // shift right n places (divides by 2ⁿ)
```

`&^` deserves a stare — most languages spell it `& ~`, but Go mints a
dedicated operator so clearing bits reads as one thought. `p &^ Write`
means "permissions minus write," no parentheses required.

## Constants that count themselves

```go
type Perm uint8

const (
    Read Perm = 1 << iota
    Write
    Execute
)
```

**`iota`** is the counter of `const` blocks: 0 in the first line, 1 in the
second, and so on. Each line repeats the *expression* `1 << iota` with its
own value, so the block unfolds to:

```go
Read    Perm = 1 << 0 // 001 == 1
Write   Perm = 1 << 1 // 010 == 2
Execute Perm = 1 << 2 // 100 == 4
```

One bit each, no overlap — three flags coexisting in a single byte. Adding
a fourth (`Delete`) is a one-line change that can't disturb the others,
which is precisely why flags beat booleans once options multiply. Combining
uses OR (`Read|Write` is `011`), and the named `Perm` type keeps these
distinct from bare numbers everywhere they travel.

## Testing a single bit

```go
// Syntax: <set> & <flag> == <flag>
func Has(p, flag Perm) bool {
    return p&flag == flag
}
```

Mask with AND, compare with the flag. `p&flag` keeps only the bits the two
share; if the result equals the flag, every requested bit was set. For
single-bit flags `!= 0` would also work — but `== flag` stays correct when
someone passes a *combined* flag like `Read|Write`, demanding both. Write
the general form now and the multi-flag case never breaks later.

Read the test once more with bit-goggles on: `Has(Read|Write, Write)` masks
`011` with `010`, gets `010`, equals the flag — true. `Has(Read, Write)`
masks `001` with `010`, gets `000` — false. Every assertion is just shared
bits surviving the mask.

## Task

Fix `Has` in `calc.go` using `&`.

## Check

```bash
go test ./exercises/02_calculator/15_bitwise/ -v
```
