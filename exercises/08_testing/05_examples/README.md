# Example Tests

Documentation rots. Every prose claim about what code does — "returns twice
its input" — decays the moment the code changes and nobody updates the
sentence. Unless the sentence *runs*. This exercise's test file contains no
`Test` function at all, yet `go test` verifies it. Here's how.

## Documentation that executes

```go
func ExampleDouble() {
    fmt.Println(Double(2))
    // Output: 4
}
```

A function named `ExampleDouble` is two things at once. To `godoc` it's
documentation, rendered beside the package's reference page as a runnable
illustration. To `go test` it's a test: the body executes, stdout is
captured, and the capture is compared against the `// Output:` comment
*line by line, byte for byte*.

The naming carries the wiring. `Example` alone demonstrates the package;
`ExampleDouble` attaches to the `Double` identifier; `ExampleDouble_suffix`
variants (lowercase suffix) show separate scenarios for one name. Rename
the function and the docs silently detach — the compiler won't complain,
but the doc page loses its illustration. Names are load-bearing here.

## Exactness is the contract

```text
// Output: 4
```

Trailing newline included, whitespace significant, order preserved. The
comparison is textual equality against captured stdout — `fmt.Println`
appends `\n`, and the expected output absorbs it. A stray space fails the
test exactly as a wrong number would. This strictness is the feature:
examples that drift from reality fail loudly instead of misleading quietly.

And that strictness flows backward into the implementation. `Double` is
deliberately trivial — `return 2 * n` — because the exercise grades the
*mechanism*, not the math. The function is a vessel; the `// Output:`
comment is the cargo. When your own packages grow examples later, keep them
this small: one behavior per example, output pinned, docs that cannot lie.

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
