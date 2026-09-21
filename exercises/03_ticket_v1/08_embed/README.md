# Embedding

Go has no inheritance. Shared fields are composed by embedding one struct
in another:

```go
type Meta struct {
    ID int
}

type Ticket struct {
    Meta
    Title string
}
```

The embedded struct's fields are promoted: `tk.ID` works directly without
writing `tk.Meta.ID`. Construction fills the embedded part explicitly:

```go
func NewTicketWithID(id int, title string) Ticket {
    // ...
}
```

## Task

Implement `NewTicketWithID` in `ticket.go` so it returns a `Ticket` with the
given id and title, and the test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/08_embed/ -v
```
