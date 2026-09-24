# If Else

After the ceremony of types and declarations, a relief: Go's `if` works the
way you'd guess. Almost. The differences from the C family are small,
deliberate, and each one removes a historical bug class.

## The shape

```go
// Syntax: if <condition> { <then> } else { <otherwise> }
func Max(a, b int) int {
    if a >= b {
        return a
    } else {
        return b
    }
}
```

Two things to notice. First, **no parentheses** around the condition —
`if (a >= b)` doesn't compile. The braces do the grouping; parentheses
would be noise. Second, **braces are mandatory**, even for one-line bodies.
Dangling-else ambiguity and the "added a second line outside the `if`"
disaster simply can't be written.

The condition must be a genuine boolean. `if x` where `x` is an `int`
fails — Go has no truthy values, no "zero means false." If that feels
pedantic, consider how many C bugs were an `=` that should have been `==`
hiding inside a tolerant condition. Pedantry is the point.

## The test already knows a trick

Read this exercise's table test once more:

```go
if got := Max(tt.a, tt.b); got != tt.want {
```

That's an `if` with a **short statement**: declare `got`, then test it, all
in the condition line. `got` lives only inside the `if`/`else` — it can't
leak into the loop and confuse the next iteration. Whenever a value exists
solely to be checked, this is the idiomatic home for it. You'll write this
shape hundreds of times before the course ends.

## Edge cases are the lesson

The table includes `{4, 4, 4}` and `{-1, -5, -1}` for a reason. Equality
(`>=` vs `>`) and negatives are where comparison code dies in production.
`Max` with `>` instead of `>=` still passes three rows and fails the tie —
which is precisely why tables list the boring cases alongside the obvious
ones. When you write your own tests later, steal this habit: equal inputs,
zero, negatives, in that order.

## Task

Fix `Max` in `calc.go` to return the larger of `a` and `b` with an `if/else`.

## Check

```bash
go test ./exercises/02_calculator/03_if_else/ -v
```
