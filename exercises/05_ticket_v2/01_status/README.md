# Status

Go has no sum-type enums, so small state sets are usually an `int`-based
type with `iota` constants:

```go
type Status int

const (
    StatusOpen Status = iota
    StatusInProgress
    StatusClosed
)
```

`iota` counts 0, 1, 2 down the const block. Because any `int` converts to
`Status`, a method can report whether a value is one of the known states:

```go
func (s Status) Valid() bool {
    // ... switch on s, true for the three known states ...
    return false
}
```

That method plays the role that matching on a simple enum plays in the Rust
course this section is adapted from: callers ask the value itself whether it
is legal instead of matching exhaustively at every use site.

## Task

Complete `Valid` in `status.go` so it returns true for `StatusOpen`,
`StatusInProgress`, and `StatusClosed`, and false for anything else.

## Check

```bash
go test ./exercises/05_ticket_v2/01_status/ -v
```
