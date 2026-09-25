# Building Strings Efficiently

Immutability has a price, and loops collect it. Each `+=` builds a brand
new string and copies everything so far — one concatenation is invisible,
ten thousand is a quadratic swamp. This exercise repeats a string `n`
times, which is exactly the shape where naive building collapses.

## Counting the copies

```go
out := ""
for i := 0; i < n; i++ {
    out += s // copy #i moves i*len(s) bytes — total O(n²)
}
```

Every iteration copies the entire accumulated result plus the new piece.
For `ConcatN("ab", 3)` nobody cares; for a thousand fragments in a request
handler, the profiler lights up. The cost hides because each line looks
cheap — the quadratic behavior is emergent, visible only in aggregate.
Whenever a loop grows a string, hear a small alarm.

## The Builder

```go
import "strings"

func ConcatN(s string, n int) string {
    var b strings.Builder
    for i := 0; i < n; i++ {
        b.WriteString(s)
    }
    return b.String()
}
```

`strings.Builder` amortizes the growth: writes accumulate in a buffer that
expands geometrically, and `String()` hands over the finished product with
one final conversion. Same loop, linear cost. `WriteString` returns an
error you'll almost always ignore (it only fails in exotic cases) —
`_, _ = b.WriteString(s)` if your linter insists, bare call if it
doesn't. The zero value is ready to use: `var b strings.Builder` needs no
constructor, the house style once more.

The zero-count case passes vacuously — loop never runs, `String()` on an
empty builder is `""`. No special case written, none needed: well-chosen
zero values keep absorbing edge cases, as they have since the ticket
section.

Could you *measure* the difference? The benchmarks lesson gave you the
tool: wrap both versions in `BenchmarkXxx` and watch ns/op diverge as `n`
grows. Performance claims without numbers are opinions; this codebase now
equips you to convert them.

## Task

Fill in `ConcatN` in `concat.go`:

```go
func ConcatN(s string, n int) string {
    // ...
}
```

Loop `n` times writing into a `strings.Builder`, then return the built string.

## Check

```bash
go test ./exercises/09_strings/05_builder/ -v
```
