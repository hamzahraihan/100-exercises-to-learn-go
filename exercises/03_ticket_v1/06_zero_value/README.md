# Zero Value

Every Go type has a zero value (`""` for strings, `0` for numbers), so a
struct is usable the moment it is declared, with no constructor call:

```go
var tk Ticket // Title == "", Description == ""
```

Well-designed structs do something sensible in that state. A helper like
`IsZero` lets callers detect an unpopulated value:

```go
func (t Ticket) IsZero() bool {
    // ...
}
```

## Task

Implement `IsZero` in `ticket.go` so it reports whether the ticket is the
zero value, and the test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/06_zero_value/ -v
```
