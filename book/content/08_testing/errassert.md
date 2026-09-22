---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Asserting Errors"
weight: 7
draft: false
---

# Asserting Errors

Go errors are values, and the test asserts both branches: the happy path
(`Divide(6, 3)` returns `2` with a nil error) and the failure path
(dividing by zero returns the sentinel). Never compare errors with `==` —
a wrapped error with context would fail that check even when it represents
the same problem. Use `errors.Is`, which unwraps the chain and matches the
sentinel no matter how it was wrapped.

## Task

Fill in `Divide` in `divide.go` so it returns `a / b`, or the existing
`ErrZeroDivisor` sentinel when `b` is 0:

```go
func Divide(a, b int) (int, error) {
	// ...
}
```

Branch on `b == 0`: one side returns the sentinel error, the other performs
the division with a nil error.

## Check

```bash
go test ./exercises/08_testing/07_errassert/ -v
```

---

*Source: `exercises/08_testing/07_errassert/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
