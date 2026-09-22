# Surviving Corrupt Files

Disk contents are untrusted input: users hand-edit files, disks glitch, and
half-written JSON happens. A loader must therefore never panic on bad data
— it returns a descriptive error instead. Good errors carry context: wrap
the underlying failure with `%w` (so callers can still unwrap it) and
include the file path, so the operator knows which file to fix. A message
without the path forces a hunt; a message with it points straight at the
problem.

## Task

Fill in `LoadTickets` in `loadtickets.go`:

```go
func LoadTickets(path string) ([]Ticket, error) {
	// ...
}
```

Read the file with `os.ReadFile`, decode it with `json.Unmarshal`, and wrap
every failure with the path (hint: `fmt.Errorf` with `%w`). The stub
returns `(nil, nil)`, so the test fails on both the good and corrupt cases.

## Check

```bash
go test ./exercises/14_persistence/03_corrupt/ -v
```
