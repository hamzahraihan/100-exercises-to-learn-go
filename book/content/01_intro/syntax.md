---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Syntax"
weight: 1
draft: false
---

# Syntax

Go functions look like this:

```go
func Compute(a, b int) int {
    return a + b
}
```

`package syntax` declares the package. `func` declares a function,
`(a, b int)` are typed parameters, the trailing `int` is the return type.

## Task

Fix `Compute` in `syntax.go` so the test passes.

## Check

```bash
go test ./exercises/01_intro/01_syntax/ -v
```

---

*Source: `exercises/01_intro/01_syntax/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
