# Importing Backups Safely

Export writes trustingly; import must not read that way. A backup file may
be truncated mid-download, hand-edited by an optimist, or minted by an
older version with looser rules. Every byte of it re-enters your program
through the front door — so import validates like a border guard, not a
librarian.

## Distrust, then admit

```go
func Import(path string) ([]Ticket, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("read %s: %w", path, err)
    }
    var ts []Ticket
    if err := json.Unmarshal(data, &ts); err != nil {
        return nil, fmt.Errorf("decode %s: %w", path, err)
    }
    for _, t := range ts {
        if t.Title == "" {
            return nil, fmt.Errorf("import %s: ticket %d has empty title", path, t.ID)
        }
    }
    return ts, nil
}
```

Three gates in sequence: readable, decodable, *acceptable*. The first two
mirror the corrupt-file lesson (path-carrying wraps, `%w` chains intact).
The third is new — domain validation over every element. An empty title
sailed through decoding without complaint (decoders check shapes, not
rules), so the importer checks the rules itself: *every* ticket, not just
the first offender found. One bad apple rejects the shipment.

## All or nothing

Note what happens on failure: `nil, err` — the whole import refused, no
partial state admitted. This is **replace semantics**: validated imports
become the *entire* state, and failed imports change nothing. The
alternative — merging good records while skipping bad ones — leaves the
store in a state no input ever described, half old and half new, with the
rejected records' absence unexplained. Replace is auditable: before or
after, never between.

The test pins both sides: the good file imports whole (one ticket, title
intact), the empty-titled file earns its rejection. Validation errors name
the offense; which ticket (`t.ID`) joins the message so operators fix the
file instead of hunting it — errors that locate, the corrupt-file
discipline, extended from files to records.

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
