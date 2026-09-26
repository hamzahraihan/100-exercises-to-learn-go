# Saving Tickets to Disk

Streams and scanners handle flow; sometimes you just want the whole file,
right now, in memory. The `os` package reduces that to two calls — and
this exercise pairs them into the oldest reliable pattern in storage: save,
then load back, then compare.

## Disk in two calls

```go
// Save: path, bytes, permissions
err := os.WriteFile(path, data, 0o644)

// Load: path back into bytes
got, err := os.ReadFile(path)
```

`os.WriteFile` creates or truncates-then-writes in one step; `os.ReadFile`
slurps the entire content back. Small configs, JSON tickets, test fixtures
— whole-file operations cover an enormous share of real file I/O, and both
directions fit on one line each. (`LoadTicket` here is already written;
your half is the save.)

## Reading the permission bits

```go
0o644 // owner read+write, group read, others read
```

The `0o` prefix marks octal — permissions are three bit-groups, and octal
digits map to groups cleanly. `6` is read+write, `4` read-only: owner gets
`6`, everyone else `4`. `0o644` is the standard for data files (readable
by all, writable by owner); executables and private keys want tighter
masks (`0o755`, `0o600`), and reaching for those should be deliberate, not
habitual. One literal, and the file's social contract is set at creation.

## Scratch space, not repo space

```go
dir := t.TempDir()
path := filepath.Join(dir, "ticket.json")
```

Tests that touch the filesystem must never write into the repo — stray
files pollute checkouts, collide in parallel runs, and linger after
failures. `t.TempDir()` returns a fresh directory per test, removed
automatically during cleanup, unique even under `-parallel`. `filepath.Join`
assembles the path with the OS-correct separator instead of hardcoded
slashes. Together: hermetic file tests, zero residue, portable across
platforms. The JSON-log exercise next door leans on the same pair — learn
the choreography once, reuse it everywhere files meet tests.

The test closes the loop the only convincing way: save known bytes, load
them back, demand byte equality. Round-trip assertions again — the
marshaling lesson's shape, now with a filesystem in the middle.

## Task

Fill in `SaveTicket` in `ticketfile.go`:

```go
func SaveTicket(path string, data []byte) error {
    // ...
}
```

Use `os.WriteFile` with `0o644` permissions to persist `data` at `path`.

## Check

```bash
go test ./exercises/12_io/04_file/ -v
```
