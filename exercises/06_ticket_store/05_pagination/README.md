# Pagination

Paging slices a window out of a list: `offset` says where the window starts
and `limit` says how many elements it may hold. Both arrive from callers you
do not control, so clamp them into range first — negative values, an offset
past the end, or a non-positive limit must all yield an empty result instead
of panicking on an out-of-range slice expression. Note the test asserts
emptiness with `len(got) != 0`, so any empty slice counts: `nil` and
zero-length non-nil slices are equally valid, and your implementation need
not prefer one over the other. This is the Go equivalent of the paging step
in the ticket-store section of the Rust course this section is adapted from.

## Task

Complete `Page` in `store.go` so it returns up to `limit` tickets starting
at `offset`, clamps out-of-range and non-positive inputs to an empty (len 0)
result without panicking, and the test passes:

```go
func Page(ts []Ticket, offset, limit int) []Ticket {
	// ...
}
```

## Check

```bash
go test ./exercises/06_ticket_store/05_pagination/ -v
```
