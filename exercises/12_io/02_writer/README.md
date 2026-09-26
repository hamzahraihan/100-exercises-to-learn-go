# Writing Lines

Readers drain; writers fill. The mirror interface takes your bytes and
reports what it accepted — and the gap between "accepted" and "all of it"
is where careful writers earn their keep.

## The mirror contract

```go
// The entire io.Writer contract:
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

`Write` accepts up to `len(p)` bytes and reports how many landed. A
`nil` error with `n < len(p)` is legal — short writes happen on sockets
and pipes — which is why robust code loops or delegates rather than
assuming full acceptance. For this exercise the writes are small and local,
so one call per line suffices; the *shape* still returns the first error
met:

```go
func WriteLines(w io.Writer, lines []string) error {
    for _, s := range lines {
        if _, err := fmt.Fprintf(w, "%s\n", s); err != nil {
            return err
        }
    }
    return nil
}
```

`fmt.Fprintf` formats straight into the writer — `%s` for the line, a
literal `\n` the test demands after each. First error aborts the loop and
travels to the caller untouched; success returns the zero error, `nil`.
Partial output plus an error is the honest outcome: the caller learns how
far it got (by counting) and why it stopped (by reading).

## The buffer double

```go
var buf bytes.Buffer
WriteLines(&buf, []string{"a", "b"})
buf.String() // "a\nb\n"
```

Tests capture output without files: `bytes.Buffer` implements `Writer` in
memory, and `&buf` passes the address (writes mutate, so the pointer —
the setters lesson, at container scale). Assert on `buf.String()` and the
test inspects exactly what a file or socket would have received. Same
function, production sink or test double, zero changes: the interface
payoff again, now flowing outward.

Real programs often interpose a buffered layer (`bufio.Writer`) between
logic and destination, flushing once at the end instead of paying per
small write. Today's writes go direct — the buffering lesson arrives when
flush discipline earns its own exercise.

## Task

Fill in `WriteLines` in `writelines.go`:

```go
func WriteLines(w io.Writer, lines []string) error {
    // ...
}
```

Range over `lines`, writing each one plus a `"\n"`, and return the first
error you meet (or nil when all writes succeed).

## Check

```bash
go test ./exercises/12_io/02_writer/ -v
```
