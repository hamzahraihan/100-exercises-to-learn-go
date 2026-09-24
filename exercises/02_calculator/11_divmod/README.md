# Div Mod

The syntax lesson promised that multiple return values are normal in Go.
This exercise cashes that promise in: division produces *two* answers, and
Go hands you both at once.

## One call, two results

```go
// Syntax: return <first>, <second>
func DivMod(a, b int) (quotient, remainder int) {
    return a / b, a % b
}
```

The caller unpacks them the same way:

```go
q, r := DivMod(7, 3) // q == 2, r == 1
```

No output parameters, no tuple type, no result struct for a two-value
answer. When a computation naturally yields a pair — quotient and
remainder, value and flag, result and error — the signature says so
directly. You've already used this shape three times over: `(sum,
overflow)`, `(value, ok)`-style checks, and now `(quotient, remainder)`.
Same rhythm, new verse.

## Truncation and its shadow

Integer `/` truncates toward zero: `7/3` is 2, and `-7/3` is -2 (not -3).
The remainder `%` keeps the dividend's sign: `-7%3` is -1. Together they
obey one invariant worth memorizing:

```go
// Always true (for b != 0): a == (a/b)*b + a%b
```

If your mental model says "remainder is always positive," negatives will
correct you at the worst hour. When in doubt, reconstruct: quotient times
divisor plus remainder must rebuild the dividend.

And `b == 0`? Integer division by zero panics — the emergency brake from
the panics lesson, built into the operator itself. `DivMod` inherits that
behavior for free: no guard needed, because crashing on a violated
contract *is* the contract. If this function ever needs to survive zero,
its signature would grow an error return — a transformation the ticket
exercises will rehearse until it's reflex.

## Task

Fix `DivMod` in `calc.go` to return `a / b` and `a % b`.

## Check

```bash
go test ./exercises/02_calculator/11_divmod/ -v
```
