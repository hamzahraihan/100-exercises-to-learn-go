# Ticket Logs (JSON Lines)

A ticket store that forgets on restart is a demo, not a system. This
exercise — the section's capstone — persists tickets durably in the
simplest format that works: **JSON lines**, one object per line, appended
forever and read back in order.

## The format that appends

```json
{"id":1,"title":"a"}
{"id":2,"title":"b"}
```

Each line is a complete, independent JSON document. Writers append without
reading or rewriting anything; readers decode incrementally, line by line.
Corruption in one line can't cascade into its neighbors; `grep` works on
the log; humans can read it raw. Compared against a single JSON array,
JSON lines trades pretty-printing for appendability — and for logs,
appendability wins every argument.

## Opening for append

```go
f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
if err != nil {
    return err
}
defer f.Close()
```

Three flags, each earning its place: `O_APPEND` positions every write at
the end (no manual seeking, no clobbering), `O_CREATE` conjures the file
on first use, `O_WRONLY` declares write-only intent. The permission bits
from the file lesson ride along. And `defer f.Close()` releases the handle
when the function exits — resource cleanup via defer, the same reflex as
mutex unlocking, now guarding a file descriptor instead of a lock.

Then the encoder does the line work for you:

```go
return json.NewEncoder(f).Encode(t)
```

`json.Encoder.Encode` marshals *and appends the newline* — the line
structure of the whole format, delivered by one method call. Forgetting
the newline would fuse two objects into invalid JSON; the encoder makes
that failure unrepresentable. Read the already-written `ReadLog` as the
mirror: scanner splitting lines (the scanner lesson, in production dress),
`Unmarshal` per line, `append` into order, `sc.Err()` checked at the end.
Every io lesson this section taught, composed into one function pair.

## The test's choreography

Append twice, read back, demand both tickets in order with the right ids.
`filepath.Join(t.TempDir(), "tickets.jsonl")` keeps the filesystem
hermetic — the scratch discipline from the file lesson, unchanged. The
`.jsonl` extension advertises the format to future readers; conventions
like this cost nothing and document everything.

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
