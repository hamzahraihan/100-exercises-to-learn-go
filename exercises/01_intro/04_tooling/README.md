# gofmt and go vet

Go ends a debate most languages never settle: *how should code look?* The
answer is a tool, not an opinion. This exercise's stub is Exhibit A — open
`tooling.go` and look at the indentation before you fix anything.

```go
func Quote(s string) string {
return s
}
```

That flush-left `return` compiles perfectly. It's also wrong — not
logically, but socially. Every Go programmer's eyes expect a tab there, and
the toolchain enforces the expectation.

## `gofmt`: formatting without arguments

`gofmt` rewrites your files into the one canonical layout: tabs for
indentation, fixed brace placement, aligned comments. No config file, no
style guide to memorize, no pull-request threads about spacing.

Two invocations cover daily life:

```bash
gofmt -l .   # list files that differ from standard (the "lint" mode)
gofmt -w tooling.go  # rewrite the file in place
```

Run the first on this exercise and `tooling.go` shows up — the missing tab
gives it away. Run the second and the file snaps into shape. Make `-l .`
returning nothing a personal definition of done; the rest of the course
assumes formatted code.

Why tabs and not spaces? So every editor can *render* indentation at the
width its human prefers while the *file* stays identical. The format is
fixed, the display is yours. That separation is the whole philosophy:
mechanical decisions belong to machines.

## `go vet`: suspicion, automated

If `gofmt` handles how code looks, `go vet` handles how code *smells*. It
scans for patterns that compile yet are almost certainly mistakes — the
classic being a `Printf` verb that doesn't match its argument:

```go
// Compiles. Vet flags it: %d fed a string.
fmt.Printf("%d", "not a number")
```

At runtime that prints `%!d(string=not a number)` — garbage where a number
should be. Vet catches it statically, before any user does. Its catalogue
covers unreachable code, suspicious mutex copies, misplaced loop variables,
and more; each check exists because someone, somewhere, shipped that exact
bug.

The course rule from the welcome page now earns its keep:

```bash
go vet ./...
go build ./...
```

Both must pass with *unsolved* exercises. Tests going red means your logic
is wrong; vet or build going red means your program is broken. Different
signals, different fixes — learn to tell them apart and debugging gets
twice as fast.

## Task

Fix `Quote` in `tooling.go`, then run `gofmt -w tooling.go`,
`go vet ./exercises/01_intro/04_tooling/` and the check below.

## Check

```bash
go test ./exercises/01_intro/04_tooling/ -v
```
