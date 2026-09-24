# Panics

Some errors shouldn't be handled. They should stop everything, loudly, with
a stack trace pointing at the crime scene. Go's word for that is `panic` —
and dividing by zero is its classic stage.

## The emergency brake

```go
// Syntax: panic(<value>)
func Divide(a, b int) int {
    if b == 0 {
        panic("division by zero")
    }
    return a / b
}
```

`panic` halts normal execution immediately and unwinds the stack, running
deferred cleanup on the way out. Unrecovered, it crashes the program. That
sounds violent because it is — which is why panics are reserved for
**programmer errors**: broken invariants, impossible states, misuse of an
API. A zero divisor passed to integer division is exactly that: the caller
violated the contract, and pretending otherwise would compute garbage with
a straight face.

Contrast this with the errors coming in the ticket section. A user typing a
bad title is *expected* — bad input happens daily and deserves a returned
error the caller can display. A zero divisor in `Divide` means the program
itself is wrong. Panic for bugs, return errors for conditions. Mixing the
two up is one of the most common Go design mistakes, so the course makes
you practice both sides of the line.

## Reading the test's safety net

The second test does something you've never seen:

```go
func TestDivideByZeroPanics(t *testing.T) {
    defer func() {
        if recover() == nil {
            t.Fatal("Divide(1, 0) did not panic")
        }
    }()
    _ = Divide(1, 0)
}
```

`recover()` catches a panic and hands back the value it carried
(`"division by zero"` here) — or `nil` if nothing panicked. Wrapped in a
deferred function, it turns a crash into an observation: the test *expects*
the explosion and fails only if the explosion never comes. You don't need
to write this pattern yourself yet. Just recognize it: `defer` + `recover`
is how Go tests assert that code refuses to proceed.

One more detail: `_ = Divide(1, 0)`. The blank identifier discards the
return value the test doesn't care about — the same `_` that skipped loop
indexes in the workflow lesson. Discarding is deliberate; Go forbids
silently ignoring values any other way.

## Task

Fix `Divide` in `calc.go` to `panic("division by zero")` when `b == 0`,
otherwise return `a / b`.

## Check

```bash
go test ./exercises/02_calculator/04_panics/ -v
```
