# Reading Everything

An `io.Reader` is a stream you drain by reading until you hit `EOF`: each
`Read` call fills your buffer and reports how many bytes it took, and the
final call reports `io.EOF` to say the stream is done. The `io` package
ships a helper that runs that loop for you. In tests, a `strings.Reader`
(or a `bytes.Buffer`) stands in for files and network connections, so you
can exercise reader code without touching disk. The stub returns `""` for
everything, so the test fails on the assertion.

## Task

Fill in `ReadAll` in `readall.go`:

```go
func ReadAll(r io.Reader) (string, error) {
	// ...
}
```

Use `io.ReadAll` to drain `r` to the end and return the bytes as a string.

## Check

```bash
go test ./exercises/12_io/01_reader/ -v
```
