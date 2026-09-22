---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Zero Value"
weight: 6
draft: false
---

# Zero Value

Every Go type has a zero value (`""` for strings, `0` for numbers), so a
struct is usable the moment it is declared, with no constructor call:

```go
var tk Ticket // Title == "", Description == ""
```

Well-designed structs do something sensible in that state. A helper like
`IsZero` lets callers detect an unpopulated value:

```go
func (t Ticket) IsZero() bool {
    // ...
}
```

## Task

Implement `IsZero` in `ticket.go` so it reports whether the ticket is the
zero value, and the test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/06_zero_value/ -v
```

---

*Source: `exercises/03_ticket_v1/06_zero_value/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
