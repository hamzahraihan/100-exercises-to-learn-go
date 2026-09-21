# Embedding Interfaces

Small interfaces compose: one interface can embed others instead of
relisting every method. The standard library does this with
`io.ReadWriter`, and this exercise mirrors that shape:

```go
type ReadWriter interface {
    Reader
    Writer
}
```

A type satisfies `ReadWriter` by implementing all the embedded methods.
The `var _ ReadWriter = (*Bucket)(nil)` line is a compile-time check:
it fails to build if `*Bucket` ever stops implementing the interface,
long before any test runs.

## Task

Implement `Write` and `Read` on `*Bucket` in `bucket.go` so the stored
string round-trips and the test passes.

## Check

```bash
go test ./exercises/04_interfaces/05_embed_iface/ -v
```
