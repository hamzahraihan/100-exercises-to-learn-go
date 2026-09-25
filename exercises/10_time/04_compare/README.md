# Testable Time

`IsExpired` looks innocent: compare now against the expiry. But "now" is
the enemy of testing — every run reads a different clock, so no assertion
can pin the answer. Code that calls `time.Now` directly is untestable by
construction. This exercise cuts the wire between the logic and the clock.

## The seam: a variable holding a function

```go
// Now is the clock. Tests replace it with a fixed time; production
// leaves it as time.Now.
var Now = time.Now
```

Functions are values in Go, so a package-level variable can *be* the
clock. Production never touches it — `Now` starts as `time.Now` and stays
that way. Tests reassign it to a fixture:

```go
old := Now
Now = func() time.Time { return time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC) }
defer func() { Now = old }()
```

Three moves, each load-bearing. Save the original (`old := Now`), install
the fixture (January 2nd, forever), and `defer` the restoration so later
tests inherit a working clock even if this one fails mid-way. Forgetting
the restore poisons every subsequent test in the package with a frozen
2026 — the `defer` isn't tidiness, it's quarantine.

## Logic reads the seam, never the wall

```go
func IsExpired(expiry time.Time) bool {
    return Now().After(expiry)
}
```

The function calls `Now()`, not `time.Now()`. One indirection, total
testability: under the fixture, January 1st is expired and January 3rd is
not, deterministically, on every machine, forever. The test asserts both
sides of the boundary — past *and* future — because a function returning
constant `true` would pass half the suite. Boundaries get both neighbors;
the table habit, applied to time.

And the comparator is `After`, per the parse lesson's rule: `Before`,
`After`, `Equal` for moments, `==` never. Times carry monotonic readings
and locations that `==` compares and `Equal` ignores — two fixtures for
the same instant built different ways must still compare equal, and only
the methods guarantee it.

## The general shape

This is dependency injection without frameworks: depend on a variable,
not a hardcoded call, and tests supply the double. Clocks are the classic
case (every codebase has one), but the seam works for randomness, IDs, and
network boundaries too. Whenever a test would need to control the
uncontrollable, look for the direct call and ask what a variable would buy.
Usually: everything.

## Task

Fill in `IsExpired` in `expiry.go` (keep the `Now` variable as the clock):

```go
func IsExpired(expiry time.Time) bool {
    // ...
}
```

Compare the fake clock's reading against `expiry` — expiry counts as expired
once the current time is past it.

## Check

```bash
go test ./exercises/10_time/04_compare/ -v
```
