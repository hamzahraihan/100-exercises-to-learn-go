# Formatting Time

Go formats times with a reference-time layout instead of format codes: you
write out the exact moment `Mon Jan 2 15:04:05 MST 2006` in the shape you
want, and Go maps each part (month, day, hour, and so on) to the real value.
For machine-readable timestamps the `time` package also ships constants like
`time.RFC3339`. The stub returns `""` for everything, so the test fails on
the assertion.

## Task

Fill in `FormatDay` in `day.go`:

```go
func FormatDay(t time.Time) string {
	// ...
}
```

Use `t.Format` with the `2006-01-02` layout to render just the calendar date.

## Check

```bash
go test ./exercises/10_time/01_format/ -v
```
