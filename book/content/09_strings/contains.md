---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Substring Checks"
weight: 3
draft: false
---

# Substring Checks

Often you only need to ask where a string sits inside another: does it contain
this piece, start with that prefix, end with that suffix? The `strings`
package names one function for each — `Contains`, `HasPrefix`, `HasSuffix` —
so the call reads like the question. The check here is deliberately naive (a
real validator would do much more), and the stub always says `false`, so the
table test fails on the positive rows.

## Task

Fill in `IsEmail` in `mail.go`:

```go
func IsEmail(s string) bool {
	// ...
}
```

Use `strings.Contains` to report whether the address holds an `"@"`.

## Check

```bash
go test ./exercises/09_strings/03_contains/ -v
```

---

*Source: `exercises/09_strings/03_contains/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
