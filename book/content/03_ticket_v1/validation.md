---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Validation"
weight: 2
draft: false
---

# Validation

Methods on a struct use a receiver:

```go
func (t Ticket) Validate() error {
    // return nil when valid, an error otherwise
}
```

The `error` type is a built-in interface; the standard way to build one
is `errors.New` or `fmt.Errorf`. A common convention is to return `nil`
for success and a non-nil error describing what is wrong.

## Task

Implement `Validate` in `ticket.go`: return an error when `Title` is empty
or `Description` is shorter than 10 characters, and `nil` otherwise, so the
test passes.

## Check

```bash
go test ./exercises/03_ticket_v1/02_validation/ -v
```

---

*Source: `exercises/03_ticket_v1/02_validation/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
