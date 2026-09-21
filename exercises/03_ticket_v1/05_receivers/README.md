# Value Receivers

A method with a value receiver works on a copy of the struct:

```go
func (t Ticket) Renamed(newTitle string) Ticket {
    // ...
}
```

The original is never modified, so the method returns the changed copy and
the caller uses the return value. This copy-out style suits small
transformations where the caller should keep the old value untouched.

## Task

Implement `Renamed` in `ticket.go` so it returns a copy of the ticket with
the new title while leaving the original unchanged, and the test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/05_receivers/ -v
```
