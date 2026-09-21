# Panic vs Error

Go convention reserves `panic` for truly unrecoverable bugs; ordinary
failures are returned as `error` values. A `Must` helper sits at the
boundary: it calls code that may panic and converts the panic into a
fallback value using `defer` plus `recover`:

```go
func MustParse(s string) (n int) {
    // ... defer a func that recovers and sets n to a fallback ...
    return parseOrPanic(s)
}
```

The deferred function runs even while a panic unwinds, `recover` stops the
unwinding, and the named return lets the deferred function choose the value
the caller sees — the Go equivalent of catching an unrecoverable failure at
a safe boundary in the Rust course this section is adapted from.

## Task

Complete `MustParse` in `must.go` so it returns `len(s)`, or `-1` when `s`
is empty, with no panic escaping to the caller.

## Check

```bash
go test ./exercises/05_ticket_v2/05_panic_vs_error/ -v
```
