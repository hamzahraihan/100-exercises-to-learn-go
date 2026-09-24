# Constructor

The previous exercise taught tickets to judge themselves — after they
already exist. That's backwards. By the time `Validate` complains, the
invalid ticket has been passed around, stored, maybe rendered. This
exercise moves the checkpoint to the only place it can't be skipped: the
moment of creation.

## Reject at birth

The convention you met in the structs lesson grows up here. `NewTicket`
stops returning a bare `Ticket` and starts returning a verdict:

```go
// Syntax: value and verdict, together
func NewTicket(title, description string) (Ticket, error) {
    // ...
}
```

`(Ticket, error)` is Go's most common signature shape — the paired-returns
rhythm from the divmod lesson, now carrying success-or-explanation. Valid
input comes back as `(Ticket{...}, nil)`. Invalid input comes back as
`(Ticket{}, someError)`: a zero value plus the reason, never a half-built
ticket pretending to be fine.

The rules are the ones you just implemented: non-empty title, description
of at least 10 characters. The difference is *where* they live. Validation
inside the constructor means every `Ticket` in the program was acceptable
at birth — the rest of the code can trust what it receives instead of
re-checking at every door.

## The caller's half of the deal

A two-value return demands a two-variable reception, and Go's most-written
three lines:

```go
tk, err := NewTicket("Fix bug", "Crash on login screen")
if err != nil {
    // handle it: log, wrap, return it upward — anything but ignore
}
```

`if err != nil` after every fallible call is the heartbeat of Go programs.
You'll write it hundreds of times before this course ends; uniformity is
the feature. Anyone reading your code knows exactly where failure is
handled, because it's handled the same way everywhere.

The tests show the mirror image — deliberate ignoring:

```go
if _, err := NewTicket("", "long enough description"); err == nil {
```

`_` takes the ticket nobody wants (it's a zero value anyway) while `err`
gets interrogated. Ignoring requires the blank identifier's explicit
consent; there's no accidental way to drop a value in Go.

## Task

Implement `NewTicket` in `ticket.go`: accept a non-empty title and a
description of at least 10 characters, returning the `Ticket` with a `nil`
error when valid and a descriptive error otherwise, so the test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/03_constructor/ -v
```
