# Sum To

The previous exercise taught `for` with all three clauses. Now meet the
shape Go uses *instead of* `while` — because there is no `while` keyword to
reach for.

## The loop with one clause

```go
// Syntax: for <condition> { <body> }
for n > 0 {
    // runs until n reaches 0
}
```

Drop the init and post clauses and `for` becomes a condition loop: keep
going while the condition holds. It's the same keyword doing a different
job, which means one fewer concept to learn and one fewer keyword to
misremember at 2 AM.

That said, this exercise sums `1 + 2 + ... + n`, and the counting form fits
better:

```go
func SumTo(n int) int {
    total := 0
    for i := 1; i <= n; i++ {
        total += i
    }
    return total
}
```

Same accumulator pattern as factorial, different seed: sums start at 0
(the empty sum), products at 1 (the empty product). The seed always answers
"what should this return when the loop runs zero times?" — ask that
question first and half of all loop bugs never get written.

## Guarding the entrance

The spec adds a wrinkle: `SumTo` returns 0 for `n <= 0`. With the loop
above, that happens *by accident* — `i <= n` is false immediately, the body
never runs, 0 comes back. Accidentally-correct code rots the moment someone
"improves" it, so say what you mean up front:

```go
func SumTo(n int) int {
    if n <= 0 {
        return 0
    }
    total := 0
    for i := 1; i <= n; i++ {
        total += i
    }
    return total
}
```

This is a **guard clause**: invalid input bounced at the door, main logic
unindented and unbothered. Go favors early returns over nested pyramids —
you'll see the style harden into dogma by the HTTP section. The test's
`SumTo(0)` row pins the behavior so the guard can't be "cleaned up" away.

## Task

Fix `SumTo` in `calc.go` to return `1 + 2 + ... + n` (0 for `n <= 0`).

## Check

```bash
go test ./exercises/02_calculator/06_sum_to/ -v
```
