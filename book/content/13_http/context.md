---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Contexts and Cancellation"
weight: 8
draft: false
---

# Contexts and Cancellation

Every request carries a `context.Context` that signals cancellation: when
the client disconnects or a timeout fires, its `Done()` channel closes.
Work that might outlive the request should `select` on `ctx.Done()` against
the real work, returning `ctx.Err()` when cancellation wins. Ignoring the
context means doing useless work — or blocking — for a client that is gone.

## Task

Fill in `Greet` in `greet.go`:

```go
func Greet(ctx context.Context, name string) (string, error) {
	// ...
}
```

Return `"hi "+name` after the simulated 10ms lookup, but return
`ctx.Err()` promptly if `ctx` is canceled first. The stub never watches the
context, so the test fails on the canceled case.

## Check

```bash
go test ./exercises/13_http/08_context/ -v
```

---

*Source: `exercises/13_http/08_context/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
