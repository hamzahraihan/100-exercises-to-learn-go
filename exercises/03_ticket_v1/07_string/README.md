# String

The `fmt` package checks for a `String() string` method and uses it for
printing:

```go
func (t Ticket) String() string {
    // ...
}
```

There is no `implements` keyword in Go: defining the method is enough for
`Ticket` to satisfy the `fmt.Stringer` interface implicitly. Any call such
as `fmt.Sprint(tk)` or `fmt.Println(tk)` then uses the custom text.

## Task

Implement `String` in `ticket.go` so the printed form includes the ticket's
title (the `fmt` package is available for formatting), and the test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/07_string/ -v
```
