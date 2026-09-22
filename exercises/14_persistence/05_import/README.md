# Importing Backups Safely

Import is where bad data tries to get back in: a backup may be truncated,
hand-edited, or from an older version. So every import validates before it
trusts — here, each ticket must have a non-empty title, and the whole
import fails if any ticket does not. When validation passes, this course
uses replace semantics: the imported tickets become the entire state,
rather than merging with whatever is already loaded. Replace is simpler to
reason about (no duplicate or ordering questions); merge is a later,
deliberate feature.

## Task

Fill in `Import` in `importtk.go`:

```go
func Import(path string) ([]Ticket, error) {
	// ...
}
```

Read the file with `os.ReadFile`, decode it with `json.Unmarshal`, and
reject any ticket with an empty title. The stub returns `(nil, nil)`, so
the test fails on the good file.

## Check

```bash
go test ./exercises/14_persistence/05_import/ -v
```
