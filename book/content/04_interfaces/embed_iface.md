---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Embedding Interfaces"
weight: 5
draft: false
---

# Embedding Interfaces

Small interfaces compose: one interface can embed others instead of
relisting every method. The standard library does this with
`io.ReadWriter`, and this exercise mirrors that shape:

```go
type ReadWriter interface {
    Reader
    Writer
}
```

A type satisfies `ReadWriter` by implementing all the embedded methods.
The `var _ ReadWriter = (*Bucket)(nil)` line is a compile-time check:
it fails to build if `*Bucket` ever stops implementing the interface,
long before any test runs.

## Task

Implement `Write` and `Read` on `*Bucket` in `bucket.go` so the stored
string round-trips and the test passes.

## Check

```bash
go test ./exercises/04_interfaces/05_embed_iface/ -v
```

---

*Source: `exercises/04_interfaces/05_embed_iface/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
