---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Type Switch"
weight: 4
draft: false
---

# Type Switch

An `any` value hides its dynamic type. A type switch recovers it, with
one case per type you care about:

```go
func Describe(v any) string {
    switch v := v.(type) {
    // ... cases for the supported types ...
    default:
        return "unknown"
    }
}
```

Inside each case, `v` has that case's type, so you can format it
directly. For a single type, the comma-ok form (`s, ok := v.(string)`)
is the lighter alternative; a switch wins once you branch on two or
more types.

## Task

Complete `Describe` in `describe.go` so `int` and `string` inputs are
named and everything else returns `"unknown"`.

## Check

```bash
go test ./exercises/04_interfaces/04_type_switch/ -v
```

---

*Source: `exercises/04_interfaces/04_type_switch/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
