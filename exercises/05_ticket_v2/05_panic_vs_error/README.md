# Panic vs Error

The panics lesson drew the line: panic for bugs, return errors for
conditions. This exercise lives exactly on that line — a function that
*receives* a panic from below and hands an *error-shaped answer* upward,
so its own callers never know an explosion happened inside.

## The boundary pattern

```go
func parseOrPanic(s string) int {
    if s == "" {
        panic("empty input")
    }
    return len(s)
}
```

`parseOrPanic` panics on empty input — by its own contract, callers must
not pass `""`. But `MustParse` promises something friendlier: the length,
or `-1` for empty, with *no panic escaping*. Somebody has to stand between
those two contracts and translate. That somebody is `defer` plus `recover`:

```go
// Syntax: named return + deferred guard
func MustParse(s string) (n int) {
    defer func() {
        if recover() != nil {
            n = -1
        }
    }()
    return parseOrPanic(s)
}
```

Three mechanisms interlock. The **named return** `(n int)` gives the
deferred function a variable it can assign — plain `return len(s)` results
couldn't be touched after the fact. The **deferred closure** runs during
unwinding, while the panic is still in flight. And **`recover()`** stops
the unwinding and reports what it caught; called without a panic in flight
it returns `nil`, so the `if` distinguishes "explosion happened" from
"normal return."

Trace both paths. `"hi"`: `parseOrPanic` returns 2 into `n`, deferred func
runs, `recover()` is `nil`, `n` stays 2. `""`: panic launches, unwinding
reaches the deferred func, `recover()` catches `"empty input"`, `n` becomes
-1, the function returns normally. The test's second assertion —
`MustParse("") == -1`, "no panic must escape" — executes the explosion path
and survives it.

## When Must is right (and when it isn't)

The standard library blesses this shape: `regexp.MustCompile`,
`template.Must` — same name prefix, same promise of "panics converted at a
safe boundary." Reach for it at initialization (parse the config once, fail
fast if it's broken), in tests, and in adapters over panicking code you
don't own.

Don't reach for it to *hide* bugs in your own logic. A `Must` wrapper
around code that panics from programmer error converts a loud, local crash
into a quiet, distant `-1` — exactly the silencing panics exist to prevent.
Translate panics at boundaries you chose deliberately: input edges, library
seams, startup. Everywhere else, let the crash speak.

## Task

Complete `MustParse` in `must.go` so it returns `len(s)`, or `-1` when `s`
is empty, with no panic escaping to the caller.

## Check

```bash
go test ./exercises/05_ticket_v2/05_panic_vs_error/ -v
```
