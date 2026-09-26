# Reading Everything

Files, network connections, request bodies, decompressors — wildly
different sources, one shared shape. Go funnels them all through an
interface with a single method, and this exercise drains such a stream to
its end.

## The one-method contract

```go
// The entire io.Reader contract:
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

`Read` fills your buffer `p` and reports how many bytes it took. Call it
repeatedly and the stream empties chunk by chunk — the final call reports
`io.EOF` instead of data. Note the philosophy: end-of-stream arrives as an
*error value*, not an exception, not a boolean. `EOF` isn't failure; it's
the stream's last word, and every reader loop treats it as the exit sign
rather than a problem. Helpers like the one below swallow it silently so
you rarely touch it directly.

## The loop, pre-written

```go
func ReadAll(r io.Reader) (string, error) {
    data, err := io.ReadAll(r)
    if err != nil {
        return "", err
    }
    return string(data), nil
}
```

`io.ReadAll` runs the fill-until-EOF loop for you and hands back the
accumulated bytes. Real errors (disconnected sockets, failing disks)
propagate as `err`; clean exhaustion returns the data with `nil`. The
`string(data)` conversion at the end crosses from bytes to text — the
border-crossing habit from the strconv lesson, now at stream scale.

One caution travels with this helper: it holds *everything* in memory.
Fine for test payloads and config files; dangerous for multi-gigabyte
streams. The copy exercise two doors down exists for exactly that case —
streaming without slurping. Reach for `ReadAll` when the input fits
comfortably; reach past it when "comfortably" needs measuring.

## Doubles for the real thing

```go
got, err := ReadAll(strings.NewReader("hello"))
```

The test never touches disk. `strings.NewReader` wraps a string in a
`Reader`, and `bytes.Buffer` plays the same role with mutable contents —
both stand in for files and connections so reader code gets exercised
without fixtures or cleanup. Any function taking `io.Reader` accepts these
doubles unchanged: that's the payoff of the interface. Code against the
narrow contract, and testing becomes construction instead of arrangement.

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
