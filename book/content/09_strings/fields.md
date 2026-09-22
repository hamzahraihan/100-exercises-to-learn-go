---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Splitting and Joining"
weight: 1
draft: false
---

# Splitting and Joining

Real input rarely arrives neatly: extra spaces, tabs, and newlines surround the
words you care about. `strings.Fields` splits on runs of whitespace and drops
the empties, which is what you want for human-typed text — unlike `Split`,
which cuts on one exact separator and keeps every empty piece. Going the other
way, joining places one separator between each element; the stub returns `nil`
and `""` for everything, so both tests fail at once.

## Task

Fill in `Words` and `JoinWords` in `words.go`:

```go
func Words(s string) []string {
	// ...
}
```

```go
func JoinWords(w []string) string {
	// ...
}
```

Use `strings.Fields` to split on whitespace runs, and `strings.Join` with a
single-space separator to put words back together.

## Check

```bash
go test ./exercises/09_strings/01_fields/ -v
```

---

*Source: `exercises/09_strings/01_fields/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
