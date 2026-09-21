# Table-Driven Tests

Most Go tests follow one pattern: a slice of `{input, want}` cases, a loop
that calls the function under test, and `t.Errorf` (not `Fatalf`) on mismatch
so every case reports instead of stopping at the first failure. The test here
checks four scores against four grades; the stub returns `""` for everything,
so all four rows fail at once — which is exactly the point.

## Task

Fill in `Grade` in `grade.go` so each score maps to its letter:

```go
func Grade(score int) string {
	// ...
}
```

Use an if/else chain on the 90/80/70 cutoffs (90+ is "A", 80+ is "B", 70+
is "C", anything below is "F").

## Check

```bash
go test ./exercises/08_testing/01_table/ -v
```
