# String/Number Conversion

Programs constantly cross the boundary between text and numbers: parsing what
the user typed, formatting what to show back. `strconv.Atoi` turns a string
into an `int` (reporting an error for junk like `"x"`), and `Itoa` goes the
other direction. The stub returns `(0, nil)` no matter what, so the valid
addition fails on the sum and the invalid input fails on the missing error.

## Task

Fill in `SumStrings` in `parse.go`:

```go
func SumStrings(a, b string) (int, error) {
	// ...
}
```

Use `strconv.Atoi` on each input and return their sum, passing conversion
errors back to the caller as-is.

## Check

```bash
go test ./exercises/09_strings/06_strconv/ -v
```
