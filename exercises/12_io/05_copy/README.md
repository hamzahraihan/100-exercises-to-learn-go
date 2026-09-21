# Copying Streams

Copying between a reader and a writer never needs the whole payload in
memory: `io.Copy` streams chunks from source to destination until `EOF`,
and `io.CopyN` stops after exactly `n` bytes, reporting how many it moved.
That keeps memory flat no matter how large the input is. The stub copies
nothing, so the test fails on the byte count and buffer contents.

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
