# Wrap

A sentinel error names a failure callers can test for. Wrapping adds context
while preserving that identity, using the `%w` verb in `fmt.Errorf`:

```go
var ErrNotFound = errors.New("not found")

func FindTicket(ids []int, id int) (int, error) {
    // ... loop over ids, return the index on a match ...
    return -1, fmt.Errorf("...: %w", ErrNotFound)
}
```

Callers then use `errors.Is(err, ErrNotFound)` to detect the miss, even
though the message carries extra context — the Go equivalent of matching on
a specific error variant in the Rust course this section is adapted from.

## Task

Complete `FindTicket` in `ticket.go` so it returns the index of `id`, or a
wrapped `ErrNotFound` when the id is absent.

## Check

```bash
go test ./exercises/05_ticket_v2/03_wrap/ -v
```
