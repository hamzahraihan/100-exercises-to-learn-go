---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Channel Pipeline"
weight: 2
draft: false
---

# Channel Pipeline

A channel lets goroutines hand values back to the code that launched them:
each worker computes one result and sends it, and the caller collects the
sends into the output slice. The trick is keeping input order — send each
result tagged with its index (or write directly to `out[i]` from the
worker) and wait for every goroutine before returning. That is the Go
equivalent of the message-passing step in the Rust course this section is
adapted from.

## Task

Complete `DoubleAll` in `pipeline.go` so each input is doubled via a
goroutine and a channel, results land in input order, and the test passes.

## Check

```bash
go test ./exercises/07_concurrency/02_channels/ -v
```

---

*Source: `exercises/07_concurrency/02_channels/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
