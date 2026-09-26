# Scanning Words

Reading everything works until the input doesn't fit the plan — a
two-gigabyte log, a never-ending pipe. Then "drain it all" becomes the
bug, and the fix is processing *pieces*: lines, words, tokens, one at a
time. Doing that by hand means buffers, partial tokens, and refill logic.
`bufio.Scanner` means not doing any of it.

## Tokens without bookkeeping

```go
func WordCount(r io.Reader) (int, error) {
    sc := bufio.NewScanner(r)
    sc.Split(bufio.ScanWords)
    count := 0
    for sc.Scan() {
        count++
    }
    return count, sc.Err()
}
```

Three moves. Wrap the reader in a `Scanner` — the `bufio` package owns the
buffer from here. Choose the token shape with a **split function**:
`ScanWords` for whitespace-separated words (cousin to `strings.Fields`,
streaming), `ScanLines` the default for lines, custom functions for
anything stranger. Then `Scan()` advances token by token, `false` when done,
`Text()` (or `Bytes()`) revealing each piece.

The loop never names the tokens — counting needs only their existence.
Empty input scans zero times and returns `(0, nil)`, no special case: the
test's second assertion pins exactly that, and the shape delivers it
naturally.

## The error check that matters

```go
return count, sc.Err()
```

When `Scan` returns `false`, two histories are possible: clean exhaustion
(`EOF`) or a real failure mid-stream (disk error, broken pipe). `sc.Err()`
distinguishes them — `nil` for the clean end, the failure otherwise. Skip
this call and both histories look identical: silent truncation wearing
success's clothes. *Always* check `Err()` after a scan loop; the one time
it matters pays for every time it didn't.

Note what's absent: no `io.EOF` handling by you. The scanner consumes the
sentinel internally and reports through `Err()`. Layers doing their jobs —
`Reader` signals, `Scanner` interprets, your code counts.

## Task

Fill in `WordCount` in `wcount.go`:

```go
func WordCount(r io.Reader) (int, error) {
    // ...
}
```

Scan `r` with the words split function, count the tokens, and return the
count along with any scan error.

## Check

```bash
go test ./exercises/12_io/03_scanner/ -v
```
