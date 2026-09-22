---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Mutex Counter"
weight: 1
draft: false
---

# Mutex Counter

One hundred goroutines adding to one integer is the classic shared-state
race: `c.n += n` looks like a single step but compiles to load, add, store,
so two goroutines can overwrite each other and increments get lost. The fix
is a `sync.Mutex` on the struct — `Add` and `Value` each lock, touch `n`,
and unlock (usually with `defer`), so only one goroutine holds the count at
a time. That is the Go equivalent of the shared-counter step in the Rust
course this section is adapted from.

## Task

Guard `n` in `counter.go` with the existing `mu` field so 100 concurrent
`Add(1)` calls always leave `Value()` at 100, and the test passes.

## Check

```bash
go test -race ./exercises/07_concurrency/01_mutex/ -v
```

The `-race` flag is required here: without it the racy stub can
flake-pass, but under the race detector it reports `DATA RACE` warnings.

---

*Source: `exercises/07_concurrency/01_mutex/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
