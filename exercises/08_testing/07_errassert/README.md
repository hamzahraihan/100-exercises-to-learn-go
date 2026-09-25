# Asserting Errors

The calculator section divided two ways: `Divide` that *panics* on zero,
and — implicitly — the question of what the alternative looks like. Here
it is. Same function name, same arithmetic, opposite philosophy: failure
as a return value the caller must confront, not an explosion it might
survive.

## Both branches, or it isn't tested

```go
func TestDivide(t *testing.T) {
    if got, err := Divide(6, 3); err != nil || got != 2 {
        t.Fatalf("Divide(6, 3) = (%d, %v), want (2, nil)", got, err)
    }
    if _, err := Divide(1, 0); !errors.Is(err, ErrZeroDivisor) {
        t.Fatalf("Divide(1, 0) err = %v, want ErrZeroDivisor", err)
    }
}
```

Error-returning functions have two contracts, and the test pins both. The
happy path asserts *two facts*: the value is right **and** the error is
`nil` — checking only the quotient would pass a function that divides
correctly while crying wolf. The failure path asserts identity via
`errors.Is`, never `==`: a wrapped `ErrZeroDivisor` with context attached
represents the same problem, and `==` would blind the test to it. The
sentinel lesson's rule, enforced at the assertion site.

## The implementation shape

```go
func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, ErrZeroDivisor
    }
    return a / b, nil
}
```

Guard clause first (the oldest habit in this course), sentinel on the
failure branch, `nil` error riding alongside every success. The zero value
`0` accompanies the error — meaningless, conventional, ignored by callers
who check `err` before touching the result. Which they must: using the
value without checking is the one unforgivable sin of Go error handling,
and every linter worth its salt flags it.

Set this beside the panicking `Divide` from the calculator section and you
hold the complete decision both ways. Panic version: contract violation,
caller is buggy, crash tells them. Error version: expected condition,
caller decides, zero divisor is Tuesday. Same arithmetic, same name,
different social contract — and choosing between them deliberately, per
function, is a working Go programmer's daily bread.

## Task

Fill in `Divide` in `divide.go` so it returns `a / b`, or the existing
`ErrZeroDivisor` sentinel when `b` is 0:

```go
func Divide(a, b int) (int, error) {
    // ...
}
```

Branch on `b == 0`: one side returns the sentinel error, the other performs
the division with a nil error.

## Check

```bash
go test ./exercises/08_testing/07_errassert/ -v
```
