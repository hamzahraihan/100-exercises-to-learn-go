# Embedding Interfaces

The Stringer lesson defined one-method interfaces. Real programs need
bigger capabilities — readers *and* writers, closers *and* flushers. You
could relist every method in every combination. Or you could compose,
exactly the way embedded structs compose fields.

## Small pieces, assembled

```go
type Reader interface {
    Read() string
}

type Writer interface {
    Write(s string)
}

type ReadWriter interface {
    Reader
    Writer
}
```

`ReadWriter` embeds both: one name, two capabilities, zero relisting. The
standard library works this way throughout — `io.ReadWriter` is precisely
this shape over byte-oriented methods. Define small interfaces for single
behaviors, embed them into the combinations programs actually need, and no
method signature is ever written twice.

Satisfaction works as before, just more of it: a type is a `ReadWriter`
when it implements *every* embedded method. `Bucket` needs both `Read` and
`Write`, on pointer receivers — the setters lesson already taught you why
a storing method must mutate through `*Bucket`:

```go
func (b *Bucket) Write(s string) {
    b.data = s
}

func (b *Bucket) Read() string {
    return b.data
}
```

## The line that never runs

```go
var _ ReadWriter = (*Bucket)(nil)
```

Read it as a sentence: "a nil `*Bucket`, used as a `ReadWriter`,
assigned to nothing." It never executes — `nil` is never dereferenced,
`_` discards the result. Its entire purpose is compile-time: if `*Bucket`
ever stops implementing the interface (a method renamed, a signature
drifted), the *build* fails here, at the declaration, instead of deep in
some test three files away.

This idiom is worth adopting the moment an interface matters to you. Place
the assertion beside the type, and the compiler becomes the test you never
have to write: conformance, checked on every build, for one line. Note the
pointer: the assertion names `*Bucket` because the methods live there — a
bare `Bucket` value doesn't implement `ReadWriter`, the receiver-set rule
from the setters lesson with real teeth.

The test itself is a round trip — write `"hi"`, read `"hi"` back. Through
the round trip it proves both halves of the embedded contract in four
lines.

## Task

Implement `Write` and `Read` on `*Bucket` in `bucket.go` so the stored
string round-trips and the test passes.

## Check

```bash
go test ./exercises/04_interfaces/05_embed_iface/ -v
```
