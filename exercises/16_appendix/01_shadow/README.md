# Shadowing

The variables lesson warned you: `:=` *declares*, and in some positions it
declares trouble. This is that trouble, at last — the bug that compiles
cleanly, prints the right answer mid-function, and returns the wrong one.

## The variable that wasn't there

```go
total := 100
bonus := 20
if bonus > 0 {
    total := total + bonus
    fmt.Println(total) // 120 — looks right!
}
return total // 100 — the outer one. Bug.
```

Inside the `if` body, `total := ...` doesn't update the outer `total`. It
declares a *brand-new* variable that happens to share the name, lives
until the closing brace, then evaporates. The print inside sees 120; the
return outside sees 100. Everything looks correct everywhere you look,
except where it counts.

## The two rules of `:=`

```go
// Rule 1 — same scope: := needs ONE new variable, reuses the rest.
n, err := strconv.Atoi(a)  // both new: declares both
m, err := strconv.Atoi(b)  // m new: declares m, REUSES err
```

```go
// Rule 2 — new scope (if/for/function body): := ALWAYS declares fresh.
if bonus > 0 {
    total := total + bonus // new variable, shadows the outer total
}
```

Same operator, opposite behavior, distinguished only by scope. A new block
plus `:=` always shadows; the same block reuses whatever already exists
(provided at least one name is genuinely new, or even that fails to
compile). Neither the compiler nor `go vet` flags shadowing — it's legal
by design, for cases like scoped temporaries. Only reading catches it.

## The fixes, in order of preference

```go
if bonus > 0 {
    total += bonus // assignment, not declaration: updates the outer
}
```

Prefer `=` (or `+=`) when you mean the existing variable — say update,
write update. When you genuinely want both (a transformed local plus the
original), give them different names so future readers never wonder which
`total` a line means. And when a function accumulates across branches,
declare the accumulator once at the top and assign everywhere below: one
variable, one name, zero shadows.

## Task

Fill in `ApplyBonus` in `shadow.go` so a positive `bonus` is added to
`total`:

```go
func ApplyBonus(total, bonus int) int {
    // ...
}
```

Add with `+=` inside the `if` — a declaring `:=` there would shadow the
parameter instead of updating it.

## Check

```bash
go test ./exercises/16_appendix/01_shadow/ -v
```
