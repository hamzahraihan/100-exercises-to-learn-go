# Copying Streams

`io.ReadAll` slurps everything into memory — perfect for configs, fatal
for gigabytes. But moving bytes from a reader to a writer shouldn't
require holding them: the data passes *through*, and memory should stay
flat no matter the size. `io.Copy` is that idea, pre-written and
battle-hardened.

## Never hold it all

```go
func CopyN(dst io.Writer, src io.Reader, n int64) (int64, error) {
    return io.CopyN(dst, src, n)
}
```

`io.CopyN` shuttles exactly `n` bytes from source to destination in
chunks, returning the count moved. The test moves 5 bytes of `"hello
world"` and asserts all three outcomes — count 5, buffer `"hello"`, error
nil. Nothing larger than a chunk ever resides in memory: terabyte inputs
cost kilobytes of RAM. Its sibling `io.Copy` runs until `EOF` with no
limit; same streaming, open-ended.

Note the count type: `n int64`. Streams routinely exceed 32-bit ranges —
multi-gigabyte files are unremarkable — so the copy API counts in 64 bits
even on 32-bit platforms. When a size, offset, or count could plausibly
grow past four billion, `int64` is the working type; `int` is for things
that fit in memory address arithmetic.

## Short sources and honest errors

What if the source holds fewer than `n` bytes? `CopyN` moves what's there
and reports `io.EOF` — the end-of-stream sentinel from the reader lesson,
surfacing as a genuine error this time, because "exactly n" was the
contract and it went unmet. Callers distinguish "moved less" from "moved
all" through that error, not the count alone. Same sentinel, different
meaning by context: loop-terminator inside `ReadAll`, contract-violation
here. Read error values against their function's promise, never in
isolation.

This is also why the return carries *both* count and error. Partial
progress plus the reason it stopped — the pattern from the writer lesson
(`WriteLines` returning at the first failure) scaled to library grade.
Check the error; trust the count only when it's nil.

## Task

Fill in `CopyN` in `copyn.go`:

```go
func CopyN(dst io.Writer, src io.Reader, n int64) (int64, error) {
    // ...
}
```

Use `io.CopyN` to stream exactly `n` bytes from `src` to `dst`.

## Check

```bash
go test ./exercises/12_io/05_copy/ -v
```
