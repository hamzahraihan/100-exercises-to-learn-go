# Subtests

`t.Run` splits one test into named subtests — "above", "below", "inside" here
— and each reports separately, so a failure tells you which case broke
instead of just which function. Run with `-v` and you get a line per
subtest; `go test -run 'TestClamp/below'` runs only the "below" case, which
is handy when debugging one branch of the logic.

## Task

Fill in `Clamp` in `clamp.go` so `n` is constrained to `[lo, hi]`:

```go
func Clamp(n, lo, hi int) int {
	// ...
}
```

Return `lo` when `n` is below it, `hi` when `n` is above it, otherwise `n`
itself.

## Check

```bash
go test ./exercises/08_testing/02_subtests/ -v
```
