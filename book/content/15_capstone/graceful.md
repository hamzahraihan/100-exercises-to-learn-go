---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Graceful Shutdown"
weight: 5
draft: false
---

# Graceful Shutdown

`Close` drops connections immediately; `Shutdown` drains them: it stops
accepting new requests, waits for in-flight ones to finish, then returns —
bounded by a timeout via `context.WithTimeout` so a stuck handler cannot hang
your deploy forever. A `&http.Server{}` with no listeners shuts down at once,
which is what the test exercises.

## Task

Fill in `ShutdownGracefully` in `shutdown.go`:

```go
func ShutdownGracefully(srv *http.Server, timeout time.Duration) error {
	// ...
}
```

Derive a timeout context and pass it to `srv.Shutdown`. The stub returns an
error, so the test fails on the return value.

## Check

```bash
go test ./exercises/15_capstone/05_graceful/ -v
```

---

*Source: `exercises/15_capstone/05_graceful/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
