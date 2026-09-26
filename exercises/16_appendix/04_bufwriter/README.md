# Buffered Writes

The writer lesson sent every line straight to the destination — one call
per line, each paying the full price of a system call. For files and
sockets that price dominates: thousands of tiny writes spend more time
crossing the kernel boundary than moving bytes. The fix mirrors the
scanner lesson's reader: interpose a buffer, flush once, pay once.

## Batching the boundary

```go
func WriteLines(w io.Writer, lines []string) error {
    bw := bufio.NewWriter(w)
    for _, s := range lines {
        if _, err := bw.WriteString(s + "\n"); err != nil {
            return err
        }
    }
    return bw.Flush()
}
```

`bufio.Writer` accumulates small writes in memory and forwards them in
large chunks — the geometric-growth trick from the Builder lesson,
relocated to the I/O boundary. `WriteString` rarely fails against memory
(the error check stays regardless; habits outlive their urgency), and the
accumulated bytes sit in the buffer until told otherwise.

## Flush is the whole exercise

```go
return bw.Flush() // WITHOUT THIS: output vanishes. Test fails. Silence.
```

Buffered bytes that never flush are bytes never sent. Forget `Flush` and
the program runs clean, returns nil, and produces *nothing* — the cruelest
failure shape in I/O, success-shaped silence. The test catches it the only
way possible: asserting on the destination's contents, which stay empty
until the flush lands. Every buffered writer in every language shares this
cliff; Go's version at least returns the flush error so the fall is
catchable.

Return `Flush()`'s error rather than discarding it — the final chunk can
fail (disk full reveals itself here, not earlier), and a write function
that swallows its own ending lies about completion. `defer`ed flushes
can't easily return errors, which is why explicit flush-before-return is
the idiom: finish loudly, in order, on purpose.

## Task

Fill in `WriteLines` in `bufwriter.go` so lines accumulate in a buffer and
flush once at the end:

```go
func WriteLines(w io.Writer, lines []string) error {
    // ...
}
```

Write each line plus `"\n"` through a `bufio.Writer`, then `Flush` and
return its error.

## Check

```bash
go test ./exercises/16_appendix/04_bufwriter/ -v
```
