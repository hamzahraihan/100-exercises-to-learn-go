# Struct

A Go `struct` groups related fields into one value:

```go
type Ticket struct {
    Title       string
    Description string
}
```

There are no access modifiers in Go: a field is exported (visible outside
the package) when its name starts with a capital letter, and unexported
otherwise. That is the Go equivalent of the visibility/encapsulation lesson
in the Rust course this section is adapted from.

## Task

Populate `NewTicket` in `ticket.go` to return the struct with the given
title and description so the test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/01_struct/ -v
```
