---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Embedding"
weight: 8
draft: false
---

# Embedding

Go has no inheritance. Shared fields are composed by embedding one struct
in another:

```go
type Meta struct {
    ID int
}

type Ticket struct {
    Meta
    Title string
}
```

The embedded struct's fields are promoted: `tk.ID` works directly without
writing `tk.Meta.ID`. Construction fills the embedded part explicitly:

```go
func NewTicketWithID(id int, title string) Ticket {
    // ...
}
```

## Task

Implement `NewTicketWithID` in `ticket.go` so it returns a `Ticket` with the
given id and title, and the test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/08_embed/ -v
```

---

*Source: `exercises/03_ticket_v1/08_embed/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
