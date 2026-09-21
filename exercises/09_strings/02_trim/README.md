# Trimming

User input and filenames both carry junk on the edges: stray spaces, trailing
newlines, extensions you want gone. `strings.TrimSpace` strips all surrounding
whitespace in one call, while `strings.TrimSuffix` removes one exact trailing
piece (only if present). There is also a cutset family that trims any of a set
of characters from the ends — handy, but easy to over-trim with, so reach for
the exact-suffix form when you know the literal. Both stubs echo their input,
so both tests fail.

## Task

Fill in `Clean` and `TrimExt` in `trim.go`:

```go
func Clean(s string) string {
	// ...
}
```

```go
func TrimExt(name string) string {
	// ...
}
```

Use `strings.TrimSpace` to strip surrounding whitespace, and
`strings.TrimSuffix` to drop a trailing `".txt"`.

## Check

```bash
go test ./exercises/09_strings/02_trim/ -v
```
