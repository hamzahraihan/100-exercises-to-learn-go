# 100 Exercises to Learn Go

Self-paced Go course inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) (CC BY-NC 4.0, all prose here is original).

## Requirements

- Go 1.23+ (check with `go version`)

## How it works

Each exercise lives in `exercises/<section>/NN_name/` with three files:

- `README.md` — the concept lesson and your task
- `<name>.go` — a stub with a `// TODO:` for you to fix (it compiles, but the test fails)
- `<name>_test.go` — the test that must pass

Solve one exercise at a time:

```bash
go test ./exercises/01_intro/01_syntax/ -v
```

Check everything (expected to FAIL until you solve all exercises):

```bash
go test ./...
```

Always-safe checks (must pass even with unsolved exercises):

```bash
go vet ./...
go build ./...
```

## Solutions

Solutions live on the `solutions` branch, same paths as `main`.

## Appendix

Beyond the core 100 exercises, six bonus lessons live in
`exercises/16_appendix/`: variable shadowing, ordered merging, custom JSON
decoding, buffered writing, fuzz testing, and `strings.Cut`.

## Demo server

`cmd/tickets` runs the capstone as a real ticket API — self-contained, so
it works whether or not the exercises are solved:

```bash
go run ./cmd/tickets
curl -X POST localhost:8080/tickets -d '{"title":"Fix bug","status":"open"}'
curl localhost:8080/tickets
```

`PORT` and `DATA_FILE` configure it; state persists across restarts, and
`Ctrl-C` shuts it down gracefully.

## Book

The Hugo site under `book/` renders directly from the exercise lessons —
`exercises/*/README.md` is the single source. Before previewing or building
the site locally, regenerate the pages:

```bash
go run ./tools/bookgen
hugo --source book
```
