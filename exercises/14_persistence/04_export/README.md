# Exporting Backups

A backup is just your data serialized in a documented format so it can be
moved or restored later. Here the format is a JSON array of tickets — the
same shape the rest of the course uses. You can write it compact (one line,
smallest bytes) or pretty (indented, human-readable); both decode to the
same tickets, so pick compact for machine-made backups and pretty when a
human will read the file. Either passes, as long as it is a valid JSON
array.

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
