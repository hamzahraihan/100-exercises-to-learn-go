# Testdata Fixtures

The `testdata/` directory is a Go convention: the toolchain ignores it during
builds, so it is the standard home for files your tests read — sample inputs,
golden outputs, small poems. Relative paths in tests resolve from the
package directory, so `FirstLine("testdata/poem.txt")` works when you run
`go test` on this package but would break if the test ran elsewhere; that
locality is what makes the fixture reliable.

## Task

Fill in `FirstLine` in `poem.go` so it returns the first line of the file at
`path`:

```go
func FirstLine(path string) (string, error) {
	// ...
}
```

Read the whole file with `os.ReadFile`, then split on the first newline
(the `strings` package has a helper for splitting at most N times).

## Check

```bash
go test ./exercises/08_testing/04_testdata/ -v
```
