---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Buffered Collect"
weight: 3
draft: false
---

# Buffered Collect

An unbuffered channel blocks every send until a receiver is ready, which
couples workers to the collector's pace. A buffered channel
`make(chan int, n)` lets up to `n` sends complete without a waiting
receiver — ideal for fan-out/collect: launch `n` goroutines, each computes
one value and sends it, then the caller drains the channel into the output
slice in order. Size the buffer to the number of workers so no goroutine
blocks on send. That is the Go equivalent of the fan-out step in the Rust
course this section is adapted from.

## Task

Complete `Collect` in `collect.go` so `n` goroutines each send one doubled
index over a buffered channel and results land in order:

```go
func Collect(n int) []int {
	// ...
}
```

## Check

```bash
go test ./exercises/07_concurrency/03_buffered/ -v
```

---

*Source: `exercises/07_concurrency/03_buffered/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
