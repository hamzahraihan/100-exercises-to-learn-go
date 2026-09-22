---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Select with Timeout"
weight: 4
draft: false
---

# Select with Timeout

`select` waits on multiple channel operations and runs whichever is ready
first. Pairing the work channel against `time.After(200 * time.Millisecond)`
gives a deadline: the fast path returns `"data"`, the slow path returns a
timeout error. Tests assert the outcome (value vs error), never the elapsed
duration, and the 10ms / 500ms cases sit far from the 200ms boundary so
ordinary scheduling jitter cannot flip them. That is the Go equivalent of
the timeout step in the Rust course this section is adapted from.

## Task

Complete `Fetch` in `fetch.go` so a simulated fetch taking `d` succeeds on
the fast path and fails with a timeout error past 200ms:

```go
func Fetch(d time.Duration) (string, error) {
	// ...
}
```

## Check

```bash
go test -race ./exercises/07_concurrency/04_select/ -v
```

---

*Source: `exercises/07_concurrency/04_select/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
