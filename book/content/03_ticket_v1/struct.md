---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Struct"
weight: 1
draft: false
---

# Struct

A Go `struct` groups related fields into one value:

```go
type Ticket struct {
    Title       string
    Description string
}
```

There are no access modifiers in Go: a field is exported (visible outside
the package) when its name starts with a capital letter, and unexported
otherwise. That is the Go equivalent of the visibility/encapsulation lesson
in the Rust course this section is adapted from.

## Task

Populate `NewTicket` in `ticket.go` to return the struct with the given
title and description so the test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/01_struct/ -v
```

---

*Source: `exercises/03_ticket_v1/01_struct/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
