# Constructor

Go has no built-in constructors. The convention is a `New...` function that
builds a value and validates it before returning:

```go
func NewTicket(title, description string) (Ticket, error) {
    // ...
}
```

Putting validation inside the constructor means invalid values are rejected
at creation time, so the rest of the code can trust a `Ticket` it receives.
Return `nil` for the error when the value is valid, and a descriptive error
(built with `errors.New` or `fmt.Errorf`) otherwise.

## Task

Implement `NewTicket` in `ticket.go`: accept a non-empty title and a
description of at least 10 characters, returning the `Ticket` with a `nil`
error when valid and a descriptive error otherwise, so the test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/03_constructor/ -v
```
