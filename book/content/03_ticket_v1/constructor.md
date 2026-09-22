---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Constructor"
weight: 3
draft: false
---

# Constructor

Go has no built-in constructors. The convention is a `New...` function that
builds a value and validates it before returning:

```go
func NewTicket(title, description string) (Ticket, error) {
    // ...
}
```

Putting validation inside the constructor means invalid values are rejected
at creation time, so the rest of the code can trust a `Ticket` it receives.
Return `nil` for the error when the value is valid, and a descriptive error
(built with `errors.New` or `fmt.Errorf`) otherwise.

## Task

Implement `NewTicket` in `ticket.go`: accept a non-empty title and a
description of at least 10 characters, returning the `Ticket` with a `nil`
error when valid and a descriptive error otherwise, so the test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/03_constructor/ -v
```

---

*Source: `exercises/03_ticket_v1/03_constructor/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
