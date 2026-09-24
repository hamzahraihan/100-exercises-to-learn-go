# Factorial

Go has exactly one loop keyword: `for`. No `while`, no `do`, no `foreach`.
Every repetition in the language — counting, waiting, iterating — is some
shape of `for`. This exercise teaches the first shape: the counting loop.

## Dissecting the header

```go
// Syntax: for <init>; <condition>; <post> { <body> }
func Factorial(n int) int {
    result := 1
    for i := 1; i <= n; i++ {
        result *= i
    }
    return result
}
```

Three clauses separated by semicolons: declare the counter, state when to
keep going, say how to advance. `i++` adds one — Go has `++` and `--` as
*statements only*, so `x = i++` doesn't compile and nobody misses it.

Everything else in the snippet is pattern, not syntax. `result := 1` is the
**accumulator**: a variable that gathers the answer piece by piece. And it
starts at 1 for a reason that bites everyone once — read on.

## The two classic mistakes

**Starting the accumulator at 0.** Multiplication by zero stays zero, so
`result := 0` returns 0 for every input including ones with correct-looking
loops. The test's `{0, 1}` row exists partly to punish this: `Factorial(0)`
must be 1 (the empty product), which falls out naturally *only* when the
loop body runs zero times and the initial 1 survives. Initialize for the
empty case, not the common one.

**`<` where `<=` belongs.** `i < n` computes `(n-1)!` — off by one, the
oldest loop bug in existence. `Factorial(5)` is 120, i.e. 1·2·3·4·5, and
the 5 must be included. When a counting loop misbehaves, check the boundary
operator before anything else.

Notice what the test table does *not* include: negatives. `Factorial(-3)`
with this code returns 1 (loop never runs). Mathematically dubious, but the
specification is silent — and tests, not philosophy, define done. Later
exercises will show how to reject invalid input properly.

## Task

Fix `Factorial` in `calc.go` to return `n!` using a `for` loop.

## Check

```bash
go test ./exercises/02_calculator/05_factorial/ -v
```
