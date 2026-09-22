---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "String/Number Conversion"
weight: 6
draft: false
---

# String/Number Conversion

Programs constantly cross the boundary between text and numbers: parsing what
the user typed, formatting what to show back. `strconv.Atoi` turns a string
into an `int` (reporting an error for junk like `"x"`), and `Itoa` goes the
other direction. The stub returns `(0, nil)` no matter what, so the valid
addition fails on the sum and the invalid input fails on the missing error.

## Task

Fill in `SumStrings` in `parse.go`:

```go
func SumStrings(a, b string) (int, error) {
	// ...
}
```

Use `strconv.Atoi` on each input and return their sum, passing conversion
errors back to the caller as-is.

## Check

```bash
go test ./exercises/09_strings/06_strconv/ -v
```

---

*Source: `exercises/09_strings/06_strconv/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
