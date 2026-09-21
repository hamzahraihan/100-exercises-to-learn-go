# Example Tests

A function named `ExampleDouble` in the test file is both documentation and a
test: `godoc` renders it alongside the package docs, and `go test` runs its
body, comparing stdout against the `// Output:` comment line by line. The
name must start with `Example` followed by the identifier it demonstrates;
the output comment must match exactly, trailing newline included.

## Task

Fill in `Double` in `double.go` so it returns twice its input:

```go
func Double(n int) int {
	// ...
}
```

Then the example's `fmt.Println(Double(2))` prints `4`, matching the
`// Output: 4` line.

## Check

```bash
go test ./exercises/08_testing/05_examples/ -v
```
