# Status

Tickets need a state: open, in progress, closed. Many languages hand you an
enum — a type with exactly three inhabitants and no others. Go doesn't have
one. What it has is sneakier: an `int` wearing a nametag, plus the
discipline to treat it like an enum anyway.

## An int in a costume

```go
// Syntax: named type over int, constants counted by iota
type Status int

const (
    StatusOpen Status = iota // 0
    StatusInProgress         // 1
    StatusClosed             // 2
)
```

`Status` is a distinct type — you can't pass a bare `int` where a `Status`
is expected without converting. But the costume is thin: `iota` just counts
0, 1, 2 down the block (the bitwise lesson showed its other face), and
*any* int converts back in. `Status(99)` compiles. It means nothing, and
the compiler won't stop it.

That's the whole game of Go enums: the type gives you names and method
attachment; *you* supply the closedness the language doesn't.

## Asking the value itself

Since illegal values can exist, legality gets checked at the edges. The
idiom is a method on the type — the value interrogates itself:

```go
func (s Status) Valid() bool {
    switch s {
    case StatusOpen, StatusInProgress, StatusClosed:
        return true
    default:
        return false
    }
}
```

A `switch` with a case list reads like the membership test it is. The test
pins both sides: all three citizens valid, `Status(99)` rejected. Callers
ask `status.Valid()` instead of reimplementing the membership everywhere —
one definition, every use site consistent.

## The zero value votes

Notice who got 0: `StatusOpen`, the first constant. That means `var s
Status` — the zero value, no constructor involved — is already Open. This
is deliberate design, not accident: the zero-value lesson taught that types
should mean something harmless before initialization, and for a ticket
lifecycle, "open" is exactly the harmless start.

Order your constants accordingly. Whichever state a fresh value should
represent goes first, at zero. Future readers will assume the zero is the
default — make the assumption true.

## Task

Complete `Valid` in `status.go` so it returns true for `StatusOpen`,
`StatusInProgress`, and `StatusClosed`, and false for anything else.

## Check

```bash
go test ./exercises/05_ticket_v2/01_status/ -v
```
