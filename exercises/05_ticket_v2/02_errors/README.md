# Ticket Errors

Go signals failures with returned `error` values. Two patterns combine here:
a sentinel error callers can test for, and wrapping that adds context while
preserving the sentinel:

```go
var ErrUnknownStatus = errors.New("unknown status")

func NewTicketWithStatus(title, description string, status Status) (Ticket, error) {
    // ... validate title/description and status ...
    return Ticket{}, fmt.Errorf("...: %w", ErrUnknownStatus)
}
```

The `%w` verb wraps `ErrUnknownStatus` inside a richer message, so callers
use `errors.Is(err, ErrUnknownStatus)` to detect it — the Go equivalent of
matching on a specific error variant in the Rust course this section is
adapted from. `NewTicketWithStatus` otherwise mirrors the validation from
`03_ticket_v1`: a non-empty title and a description of at least 10
characters, plus the status must be a known `Status`.

## Task

Complete `NewTicketWithStatus` in `ticket.go` so unknown statuses wrap
`ErrUnknownStatus`, valid tickets are returned with their status set, and
the tests pass.

## Check

```bash
go test ./exercises/05_ticket_v2/02_errors/ -v
```
