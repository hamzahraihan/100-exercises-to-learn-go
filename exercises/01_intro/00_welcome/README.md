# Welcome

You've heard Go is simple. Small standard library habits, fast builds, a
toolchain that does what you mean. This course tests that reputation the
only honest way: by making you write it, one exercise at a time.

You'll go from an empty `return ""` to a working HTTP ticket API across
roughly a hundred exercises. No prior Go assumed — but you should already
know at least one programming language. Everything else, including the
unfamiliar bits, is taught when you need it.

## Learn by doing

Each exercise is a tiny broken program plus a test that describes the fix:

- `welcome.go` — a stub with a `// TODO:` for you to resolve. It compiles,
  but the test fails.
- `welcome_test.go` — the specification. Read it before touching the stub;
  it tells you exactly what "done" looks like.

```go
func TestMessage(t *testing.T) {
    got := Message()
    if got == "" {
        t.Fatal("Message() returned empty string, expected a greeting")
    }
}
```

The whole course runs on this loop: **read the test, fix the stub, run the
test**. Resist the urge to peek at solutions first — the struggle is where
the learning happens. Solutions live on the `solutions` branch for when
you're truly stuck, not for when you're merely impatient.

## What you need

- **Go 1.23+** — check with `go version`.
- A terminal. Every exercise is verified with one command:

```bash
go test ./exercises/01_intro/00_welcome/ -v
```

Two commands must *always* stay green, even with unsolved exercises —
they're your safety net throughout the course:

```bash
go vet ./...
go build ./...
```

If either breaks, you didn't break an exercise, you broke the build. Fix
that first.

## Task

Open `welcome.go` and make `Message` return a non-empty greeting.

## Check

```bash
go test ./exercises/01_intro/00_welcome/ -v
```
