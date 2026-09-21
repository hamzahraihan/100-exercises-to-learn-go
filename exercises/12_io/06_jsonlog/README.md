# Ticket Logs (JSON Lines)

A JSON-lines log stores one JSON object per line, so writers can append
without rewriting the file and readers can decode it incrementally. The
append side opens the file with `O_APPEND|O_CREATE|O_WRONLY` and encodes
each `Ticket` (`{ID int \`json:"id"\`, Title string \`json:"title"\`}`)
as a single line; the read side scans line by line and decodes each one in
order. Here `ReadLog` already works, but `AppendLog` writes nothing, so the
test fails when the read finds no file.

## Task

Fill in `AppendLog` in `ticketlog.go`:

```go
func AppendLog(path string, t Ticket) error {
	// ...
}
```

Open `path` in append mode (creating it if missing) and encode `t` as one
JSON line.

## Check

```bash
go test ./exercises/12_io/06_jsonlog/ -v
```
