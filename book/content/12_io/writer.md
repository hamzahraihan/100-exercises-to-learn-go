---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Writing Lines"
weight: 2
draft: false
---

# Writing Lines

An `io.Writer` is the mirror of a reader: you hand it bytes and it reports
how many it accepted, returning an error if something went wrong. Writers
compose well — the same code works against a file, a network connection,
or a `bytes.Buffer` in a test. Real programs often wrap writers in a
buffered layer and flush once at the end instead of paying for many small
writes. The stub writes nothing, so the test fails comparing the buffer.

## Task

Fill in `WriteLines` in `writelines.go`:

```go
func WriteLines(w io.Writer, lines []string) error {
	// ...
}
```

Range over `lines`, writing each one plus a `"\n"`, and return the first
error you meet (or nil when all writes succeed).

## Check

```bash
go test ./exercises/12_io/02_writer/ -v
```

---

*Source: `exercises/12_io/02_writer/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
