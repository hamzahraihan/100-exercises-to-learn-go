---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "gofmt and go vet"
weight: 4
draft: false
---

# gofmt and go vet

`gofmt -l .` lists files whose formatting differs from Go standard — fix
them with `gofmt -w`. `go vet ./...` catches real bugs, e.g. a `Printf`
call with mismatched verbs:

```go
// vet flags this: %d with a string argument
// fmt.Printf("%d", "not a number")
```

That snippet is only an example. Your exercise has no vet issue — vet must
stay clean while the test fails.

## Task

Fix `Quote` in `tooling.go`, then run `gofmt -w tooling.go`,
`go vet ./exercises/01_intro/04_tooling/` and the check below.

## Check

```bash
go test ./exercises/01_intro/04_tooling/ -v
```

---

*Source: `exercises/01_intro/04_tooling/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
