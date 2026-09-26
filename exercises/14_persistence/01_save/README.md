# Crash-Safe Saves

Writing a file looks atomic from the outside: one call, bytes land. It
isn't. `os.WriteFile` streams bytes over time, and a crash — power loss,
`kill -9`, a panicking sibling goroutine — can land mid-stream, leaving
half a file where the whole old data used to be. The old data is gone, the
new data is broken, and no error will ever explain what happened. This
exercise makes saves survive crashes they can't prevent.

## The disaster, precisely

```go
os.WriteFile(path, data, 0o644) // crash here → half-written file, old data gone
```

Direct writes destroy before they create: the file is truncated first,
then filled. Any interruption between those moments loses both versions.
Backups, configs, ticket stores — anything rewritten in place shares the
fate. The fix never writes the destination directly at all.

## Write aside, then swap

```go
func Save(path string, data []byte) error {
    dir := filepath.Dir(path)
    tmp, err := os.CreateTemp(dir, "tmp-*")
    if err != nil {
        return err
    }
    tmpName := tmp.Name()
    // If anything below fails, don't litter: remove the temp file.
    defer func() {
        tmp.Close()
        if err != nil {
            os.Remove(tmpName)
        }
    }()
    if _, err = tmp.Write(data); err != nil {
        return err
    }
    if err = tmp.Close(); err != nil {
        return err
    }
    return os.Rename(tmpName, path)
}
```

Four steps, each load-bearing. **Create** a temp file for staging.
**Write** the full payload into it — crashes here harm only the temp.
**Close** before renaming (on Windows an open file can't be renamed at
all; everywhere else, close flushes the last buffered bytes). **Rename**
over the target: readers see the complete old file or the complete new
one, never a mixture, because the swap itself is atomic.

## Same directory, or it isn't atomic

The temp file must live in the target's directory — `filepath.Dir(path)`
plus `os.CreateTemp` guarantees it. Only same-directory renames are atomic
on mainstream filesystems; across directories (or filesystems) a "rename"
can silently degrade into a slow copy-then-delete, reintroducing the exact
half-written window the dance exists to close. Collocating temp and target
isn't tidiness — it's the precondition the whole guarantee stands on.

And the deferred cleanup: on any failure after creation, remove the temp
so crashed saves don't accumulate litter. Note the `Close` in the defer
runs even on the success path (closing twice is harmless) — one deferred
guard covering every exit, the `defer`-as-quarantine habit from the clock
tests, now protecting a directory instead of a variable.

## Task

Fill in `Save` in `atomic.go`:

```go
func Save(path string, data []byte) error {
    // ...
}
```

Create a temp file in the target's directory with `os.CreateTemp`, write
`data` into it, `Close` it, then `os.Rename` it onto `path`. The stub
writes nothing, so the test fails when the file is missing.

## Check

```bash
go test ./exercises/14_persistence/01_save/ -v
```
