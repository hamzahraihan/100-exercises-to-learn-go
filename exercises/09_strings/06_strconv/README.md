# String/Number Conversion

Programs live on a border: humans speak text, machines compute numbers,
and every input field plus every display line is a crossing. This exercise
works the checkpoint in both directions — parse what arrived, add it, and
report honestly when the text isn't a number at all.

## Crossing with a receipt

```go
import "strconv"

n, err := strconv.Atoi("42") // 42, nil — the crossing succeeded
_, err = strconv.Atoi("x")   // 0, err  — junk gets a receipt, not a panic
```

`strconv.Atoi` (ASCII-to-integer) returns the value *and* a verdict,
paired returns doing their eternal duty. Valid text converts; junk yields
zero plus a non-nil error describing the failure. The zero is meaningless
by design — callers check `err` before touching `n`, the unforgivable-sin
rule from the error-assertion lesson, now at a border crossing.

The reverse direction is `strconv.Itoa` (integer-to-ASCII): `Itoa(42)` is
`"42"`. No error possible — every int has a spelling — so the signature
skips the verdict. Asymmetric pairs like this (`Atoi`/`Itoa`) reward
reading both directions before using either.

## Pass failures upward, unchanged

```go
func SumStrings(a, b string) (int, error) {
    x, err := strconv.Atoi(a)
    if err != nil {
        return 0, err
    }
    y, err := strconv.Atoi(b)
    if err != nil {
        return 0, err
    }
    return x + y, nil
}
```

Each conversion guarded, each failure returned as-is. No wrapping, no
rewording — at this layer there's no context worth adding ("parsing the
first operand" helps nobody), so the original error travels untouched for
the caller to interpret. Restraint is a decision: wrap where context
exists (the ticket errors lesson), pass through where it doesn't. The test
pins both halves — `("3","4")` sums to 7 with nil error, `("x","4")`
surfaces *an* error — checking presence, not prose, exactly as
`errors.Is`-style assertions should.

And the old trap, one final sighting: `string(65)` is `"A"`, not `"65"` —
conversion to a code point, not to digits. Every border crossing in Go
means what it says; `Atoi`/`Itoa` mean digits, and digits are what this
exercise trades in.

## Task

Fill in `SumStrings` in `parse.go`:

```go
func SumStrings(a, b string) (int, error) {
    // ...
}
```

Use `strconv.Atoi` on each input and return their sum, passing conversion
errors back to the caller as-is.

## Check

```bash
go test ./exercises/09_strings/06_strconv/ -v
```
