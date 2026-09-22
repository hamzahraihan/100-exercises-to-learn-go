---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Setters"
weight: 4
draft: false
---

# Setters

A method with a pointer receiver can mutate the value it is called on:

```go
func (t *Ticket) SetTitle(title string) {
    // ...
}
```

With a value receiver (`func (t Ticket) ...`), the method gets a copy, so
assignments inside are lost when the method returns. Use a pointer receiver
when the method needs to change the struct, and a value receiver when it
only reads it.

## Task

Implement `SetTitle` in `ticket.go` so it changes the ticket's title and
the test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/04_setters/ -v
```

---

*Source: `exercises/03_ticket_v1/04_setters/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
