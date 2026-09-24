# Any

Every function so far named the exact type it accepts. `SprintAny` refuses
to — and that's the whole exercise. One parameter, every possible argument.

## The box that holds anything

```go
func SprintAny(v any) string {
    return fmt.Sprint(v)
}
```

**`any`** is the alias for the empty interface, `interface{}`: the
interface with no methods, which every type therefore satisfies. A
parameter of type `any` accepts an `int`, a `string`, your `Ticket`, a
`nil` — literally anything, with no conversion and no complaint.

Inside, the value keeps its original identity. `any` doesn't erase the
type; it *hides* it behind a static type that promises nothing. `fmt.Sprint`
is built for exactly this: it inspects what arrived at runtime and formats
accordingly, which is why the same call renders `42` as `"42"` and `"hi"`
as `"hi"`.

## Freedom with the receipt removed

That freedom costs compile-time checking. With a concrete parameter, passing
the wrong type fails the build. With `any`, *everything* compiles and
mistakes surface at runtime — or worse, as silently odd output. The rule:

- Reach for `any` when you truly accept anything: formatting, logging,
  test helpers, serialization boundaries.
- Prefer a concrete type or a narrower interface the moment you know more.
  If callers pass temperatures, say `float64`. If they pass printable
  things, say `fmt.Stringer`. Each precise signature is a mistake the
  compiler catches for free.

The test pins both representatives — number and string — through the same
door. Your implementation doesn't branch (that's the *next* exercise's
job); it delegates to `fmt` and lets the runtime sort it out. One line,
total generality, eyes open about the price.

## Task

Implement `SprintAny` in `any.go` so it returns the formatted value and
the test passes.

## Check

```bash
go test ./exercises/04_interfaces/03_any/ -v
```
