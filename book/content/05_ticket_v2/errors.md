---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Ticket Errors"
weight: 2
draft: false
---

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

---

*Source: `exercises/05_ticket_v2/02_errors/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
