# Subtests

Tables collect rows; subtests give them *names*. When `TestGrade` fails you
learn a score mismatched — but which scenario? Above? Below? The boundary
itself? `t.Run` splits one test into labeled children, and each child
reports, passes, and filters independently.

## Naming the cases

```go
func TestClamp(t *testing.T) {
    t.Run("above", func(t *testing.T) {
        if got, want := Clamp(99, 0, 10), 10; got != want {
            t.Fatalf("Clamp(99, 0, 10) = %d, want %d", got, want)
        }
    })
    t.Run("below", func(t *testing.T) {
        // ...
    })
    t.Run("inside", func(t *testing.T) {
        // ...
    })
}
```

`t.Run("above", func(t *testing.T) {...})` registers a child test carrying
its own `*testing.T`. Failures name the child — `TestClamp/below` — so the
output points at the broken *scenario*, not just the broken function. With
`-v`, every subtest gets its own line, PASS or FAIL; a table's rows, by
contrast, all hide inside one verdict.

## Debugging one branch

Names are addresses. This runs only the "below" case:

```bash
go test ./exercises/08_testing/02_subtests/ -run 'TestClamp/below' -v
```

The `-run` flag takes a slash-separated path, each level a regex. Debugging
one branch of the logic means executing one branch of the logic — no
commenting out siblings, no printf archaeology. When a table grows to
twenty rows and one turns red, you'll wish they were subtests; when three
assertions share one setup, you'll wish the reverse. Tables for uniform
rows, subtests for heterogeneous stories.

## The logic under test

`Clamp` constrains `n` to `[lo, hi]` — below becomes `lo`, above becomes
`hi`, inside passes through:

```go
func Clamp(n, lo, hi int) int {
    if n < lo {
        return lo
    }
    if n > hi {
        return hi
    }
    return n
}
```

Three regions, three subtests, one assertion each. The subtests mirror the
branches exactly ("above", "below", "inside"), which is the habit to steal:
name subtests after the *behavioral region* they cover, not the inputs
they feed. Future readers learn the function's contract from the test
names alone.

## Task

Fill in `Clamp` in `clamp.go` so `n` is constrained to `[lo, hi]`:

```go
func Clamp(n, lo, hi int) int {
    // ...
}
```

Return `lo` when `n` is below it, `hi` when `n` is above it, otherwise `n`
itself.

## Check

```bash
go test ./exercises/08_testing/02_subtests/ -v
```
