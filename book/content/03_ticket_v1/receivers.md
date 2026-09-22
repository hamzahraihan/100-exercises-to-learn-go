---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Value Receivers"
weight: 5
draft: false
---

# Value Receivers

A method with a value receiver works on a copy of the struct:

```go
func (t Ticket) Renamed(newTitle string) Ticket {
    // ...
}
```

The original is never modified, so the method returns the changed copy and
the caller uses the return value. This copy-out style suits small
transformations where the caller should keep the old value untouched.

## Task

Implement `Renamed` in `ticket.go` so it returns a copy of the ticket with
the new title while leaving the original unchanged, and the test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/05_receivers/ -v
```

---

*Source: `exercises/03_ticket_v1/05_receivers/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
