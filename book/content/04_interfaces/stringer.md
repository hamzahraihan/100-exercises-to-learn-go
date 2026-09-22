---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Stringer"
weight: 1
draft: false
---

# Stringer

Many standard-library functions print values through the `fmt.Stringer`
interface:

```go
type Stringer interface {
    String() string
}
```

A Go type implements an interface implicitly, just by defining the method —
there is no `implements` declaration. Once `Order` has a `String` method,
calls like `fmt.Sprint(o)` or `fmt.Println(o)` use it automatically. That
is the Go equivalent of implementing a formatting trait in the Rust course
this section is adapted from.

## Task

Implement `String` in `order.go` so the returned string contains both the
order id and the item name, and the test passes.

## Check

```bash
go test ./exercises/04_interfaces/01_stringer/ -v
```

---

*Source: `exercises/04_interfaces/01_stringer/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
