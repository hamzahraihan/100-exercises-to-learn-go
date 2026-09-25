# Testdata Fixtures

Tests that read files need files to read — checked in, versioned, and
sitting where the test expects them. Go's answer is a directory named
`testdata/`, and the name is a contract with the toolchain, not a
suggestion.

## The directory the build ignores

Any directory called `testdata` is skipped by the Go tool: `go build`
won't compile packages inside it, `go vet` looks past it. That makes it
the standard home for everything tests consume but programs never import —
sample inputs, golden outputs, and in this exercise, a small poem:

```text
Roses are red
Violets are blue
Go tests are green
When fixtures help you
```

No special registration, no embedding directives. Name it `testdata`,
place it beside the test, and both humans and tooling understand the deal:
these files belong to the tests.

## Paths resolve from the package

```go
got, err := FirstLine("testdata/poem.txt")
```

Relative paths in tests resolve from the *package directory* — the folder
holding the test file — so this path works whenever `go test` runs this
package, on any machine, in any checkout. That's the whole reliability
story: locality instead of absolute paths, convention instead of
configuration. Run the test from the repo root, from the package dir, in
CI — same file found, same way.

## Read all, keep the head

```go
func FirstLine(path string) (string, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return "", err
    }
    return strings.SplitN(string(data), "\n", 2)[0], nil
}
```

`os.ReadFile` swallows the open/read/close dance into one call — the file
is small, so whole-content reading is the right tool. `strings.SplitN`
with `N = 2` splits at most once: head and remainder, no wasted work
splitting ninety lines you discard. Index `[0]` takes the first line;
propagating `err` untouched keeps the failure's origin intact for the
caller to report.

The test asserts both halves of the `(string, error)` contract: no error
*and* the exact line `"Roses are red"`. Error returns get tested like
values — because they are values, the errors-as-values lesson all over
again, now wearing a file path.

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
