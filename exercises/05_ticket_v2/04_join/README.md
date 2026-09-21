# Join

When several fields can be wrong at once, report every problem instead of
just the first. `errors.Join` combines per-field errors into one value that
still matches each part with `errors.Is`:

```go
var (
    ErrBadTitle = errors.New("bad title")
    ErrBadDesc = errors.New("bad description")
)

func ValidateAll(title, desc string) error {
    // ... collect one error per bad field ...
    return errors.Join(/* ... per-field errors ... */)
}
```

A valid input joins nothing and yields `nil`, so callers treat a `nil`
error as success — the Go equivalent of accumulating validation failures in
the Rust course this section is adapted from.

## Task

Complete `ValidateAll` in `ticket.go` so empty titles report `ErrBadTitle`,
short descriptions report `ErrBadDesc`, both problems are reported together,
and valid input returns `nil`.

## Check

```bash
go test ./exercises/05_ticket_v2/04_join/ -v
```
