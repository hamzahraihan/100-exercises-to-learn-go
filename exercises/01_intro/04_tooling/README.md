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
