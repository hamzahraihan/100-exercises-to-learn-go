# Parsing Time

Parsing is formatting in reverse: the layout you pass must describe the
input's exact shape, or `time.Parse` returns an error instead of a time.
That strictness is a feature — a mismatched layout fails loudly rather than
guessing. The stub returns the zero time and a nil error, so both halves of
the test fail on assertions.

## Task

Fill in `ParseDay` in `day.go`:

```go
func ParseDay(s string) (time.Time, error) {
	// ...
}
```

Use `time.Parse` with the `2006-01-02` layout; return the parse error
unchanged on bad input.

## Check

```bash
go test ./exercises/10_time/02_parse/ -v
```
