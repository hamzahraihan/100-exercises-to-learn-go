# Variables

Every value lives somewhere. Go gives you two ways to say "a box called
`x`" — and the difference between them is the difference between Go's two
audiences: the compiler that must know everything, and the human who'd
rather not repeat it.

## The two declarations

```go
var x int = 21   // long form: name, type, value — everything explicit
y := x * 2       // short form: name and value; the type is inferred
```

`var x int = 21` states the type out loud. `y := x * 2` lets the compiler
infer it — `y` is an `int` because `x * 2` is one. Same result, less ink.
The short form exists because most of the time the right-hand side already
shouts the type; writing it twice helps nobody.

But `:=` has boundaries. It works **only inside functions**, and it
*declares* — using it on an already-declared variable in the same scope is
an error (or, in some positions, an infamous rebinding surprise you'll meet
much later — not today).

## The box starts full

Declare without a value and Go fills the box itself:

```go
var total int    // 0, not garbage
var name string  // "", not null
var ok bool      // false
```

Every type has a **zero value**, and variables begin life holding it. This
is why `SumTo` and friends can declare `total` and start adding — no
uninitialized-memory roulette, ever. The struct lesson's missing fields and
this lesson's missing values are the same promise wearing different hats.

So when do you write `var`? When the type needs saying: zero-value
declarations (`var total int`), package-level variables (`:=` is banned
there), or the rare case where inference would pick a narrower type than
you want. Everywhere else, `:=` is idiomatic.

## Task

Fix `Double` in `calc.go` to return twice `x`, using `:=` to declare the result.

## Check

```bash
go test ./exercises/02_calculator/02_variables/ -v
```
