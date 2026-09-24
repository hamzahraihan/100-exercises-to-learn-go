# Zero Value

`NewTicket` promised that every ticket enters the world valid. That promise
has a hole: nobody is forced to use the constructor.

```go
var tk Ticket // compiles. No NewTicket in sight.
```

Any caller can declare a `Ticket` directly and hold… what, exactly? In most
languages, an uninitialized object is a landmine — null fields, undefined
behavior, the constructor-or-die ultimatum. Go chose differently, and this
exercise is about designing *for* that choice instead of against it.

## Useful before you're ready

Every Go type has a **zero value**, and every variable starts holding it:
`""`, `0`, `nil`, `false`, applied field by field. `var tk Ticket` is a
ticket with an empty title and empty description — fully formed, fully
safe to pass around, immediately comparable and printable.

Well-designed types *do something sensible* in that state. For `Ticket`,
"empty" is already meaningful: no title, no description, nothing to
validate yet. The constructor remains the front door for real tickets, but
the zero value is a legitimate back room — maps fill with it, decoders
decode into it, tests compare against it.

## Naming the empty state

Callers need a way to ask "is this populated?" without comparing fields
by hand. That's `IsZero`'s whole job:

```go
func (t Ticket) IsZero() bool {
    return t.Title == "" && t.Description == ""
}
```

Both fields empty means zero; either field set means real. The test
exercises both directions, including a half-filled ticket (`Title: "x"`,
empty description) that must *not* report zero — the `&&` earns its keep
there.

Notice the call syntax in the test:

```go
(Ticket{}).IsZero()
```

Methods work on composite literals, not just variables. `Ticket{}` builds
the zero value inline and interrogates it in the same breath. No
declaration, no name, no ceremony — the exact brevity the zero value was
designed for.

## Designing with zero in mind

This generalizes into a design habit: when you define a type, ask what its
zero value *means* before asking what its constructor *requires*. Ideal
answer: something harmless and detectable — an empty collection ready to
append to, a config with sane defaults, a ticket with nothing in it. Types
whose zero is useful compose effortlessly; types whose zero is poison force
every user through the constructor under threat.

`NewTicket` guards the front door. `IsZero` labels the back room. Between
them, no ticket state goes unnamed.

## Task

Implement `IsZero` in `ticket.go` so it reports whether the ticket is the
zero value, and the test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/06_zero_value/ -v
```
