# Setters

A method with a pointer receiver can mutate the value it is called on:

```go
func (t *Ticket) SetTitle(title string) {
    // ...
}
```

With a value receiver (`func (t Ticket) ...`), the method gets a copy, so
assignments inside are lost when the method returns. Use a pointer receiver
when the method needs to change the struct, and a value receiver when it
only reads it.

## Task

Implement `SetTitle` in `ticket.go` so it changes the ticket's title and
the test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/04_setters/ -v
```
