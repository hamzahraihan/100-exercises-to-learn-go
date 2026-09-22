---
# DO NOT EDIT — generated from exercises/ by tools/bookgen. Edit the source README instead.
title: "Example Tests"
weight: 5
draft: false
---

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

---

*Source: `exercises/08_testing/05_examples/README.md` · Inspired by [Mainmatter's 100 Exercises to Learn Rust](https://github.com/mainmatter/100-exercises-to-learn-rust) ([CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)) — all Go prose here is original, free for non-commercial use.*
