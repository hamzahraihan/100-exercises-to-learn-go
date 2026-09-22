---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Parsing Time"
weight: 2
draft: false
---

# Parsing Time

Parsing is formatting in reverse: the layout you pass must describe the
input's exact shape, or `time.Parse` returns an error instead of a time.
That strictness is a feature — a mismatched layout fails loudly rather than
guessing. The stub returns the zero time and a nil error, so both halves of
the test fail on assertions.

## Task

Fill in `ParseDay` in `day.go`:

```go
func ParseDay(s string) (time.Time, error) {
	// ...
}
```

Use `time.Parse` with the `2006-01-02` layout; return the parse error
unchanged on bad input.

## Check

```bash
go test ./exercises/10_time/02_parse/ -v
```

---

*Source: `exercises/10_time/02_parse/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
