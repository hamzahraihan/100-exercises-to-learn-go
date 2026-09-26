# Exporting Backups

State on disk is one copy away from oblivion — disks fail, deploys wipe,
`rm` happens. A backup is the second copy, and its value lies entirely in
being *restorable*: a documented format, written faithfully, readable by
code you haven't written yet. This exercise produces one.

## A format is a promise

```go
func Export(path string, ts []Ticket) error {
    out, err := json.Marshal(ts)
    if err != nil {
        return err
    }
    return os.WriteFile(path, out, 0o644)
}
```

Two known calls composed: `json.Marshal` (the marshaling lesson, now over
a slice — arrays encode element by element, no new machinery) plus
`os.WriteFile` with `0o644` (the file lesson's permissions, unchanged).
Export is composition, not invention — which is precisely why it belongs
this late in the course. Every prerequisite arrived earlier; the lesson is
that they're sufficient.

The format choice matters more than the code. A JSON array of tickets —
same shapes, same tags as the API — means the backup is readable by every
tool that speaks the domain: future importers, debuggers, migration
scripts, humans with `jq`. A bespoke binary blob would be smaller and
utterly opaque. Backups optimize for the reader during an incident, never
for the writer during calm.

## Pretty or compact?

```go
json.Marshal(ts)          // [{"id":1,"title":"a"},...] — smallest bytes
json.MarshalIndent(ts, "", "  ") // human-readable, larger file
```

Both decode identically, so the choice is audience, not correctness.
Machine-made backups consumed only by importers: compact, smallest disk
and wire cost. Files a human might inspect mid-incident: indented, worth
every extra byte at 3 AM. The test decodes rather than string-comparing
(the round-trip habit once more), so either passes — pick by reader, and
say which reader you picked in a comment.

## Task

Fill in `Export` in `export.go`:

```go
func Export(path string, ts []Ticket) error {
    // ...
}
```

Encode `ts` with `json.Marshal` and persist the bytes with `os.WriteFile`
using `0o644` permissions. The stub writes nothing, so the test fails when
the backup file is missing.

## Check

```bash
go test ./exercises/14_persistence/04_export/ -v
```
