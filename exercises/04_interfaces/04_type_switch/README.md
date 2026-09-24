# Type Switch

`any` accepts everything and promises nothing — so how do you *act* on what
arrived? Try calling string methods on it and the compiler stops you: as
far as the static type knows, there are no methods at all. This exercise
recovers the hidden type and branches on it.

## Opening the box

For a single suspected type, Go offers the comma-ok assertion:

```go
s, ok := v.(string) // s is a string, ok reports success
```

One type, one line. But `Describe` handles *several* types, and chained
assertions would nest into a staircase. The type switch is the multi-way
version:

```go
// Syntax: switch <typed> := <any>.(type)
func Describe(v any) string {
    switch v := v.(type) {
    case int:
        return fmt.Sprintf("int: %d", v) // v is an int here
    case string:
        return fmt.Sprintf("string: %s", v) // v is a string here
    default:
        return "unknown"
    }
}
```

The syntax looks peculiar because it is: `.(type)` works *only* in a
switch, and the `v :=` redeclares `v` inside each case with that case's
type. Within `case int`, `v` is genuinely an `int` — format it, add to it,
pass it to int-only functions. The compiler knows, because the branch
proved it.

## The cases that matter

The test table is the spec, read it as one:

```go
{42, "int: 42"},
{"hi", "string: hi"},
{1.5, "unknown"},
{nil, "unknown"},
```

Two named types with formatted output (`"int: 42"` — the label is part of
the contract, not decoration), and a `default` that absorbs everything
else: unlisted types *and* `nil`, which matches no case and falls through
to the bottom. A missing `default` would silently return `""` for floats —
compiling, running, and lying. The `default` isn't defensive decoration;
it's the branch the table's last two rows execute.

When the type list grows past a handful, reconsider: a switch with twelve
cases is an interface begging to be defined. Branch on types you *know*;
abstract over behavior you *don't*. That judgment — switch versus
interface — is one of Go's central design calls, and now you've practiced
one side of it.

## Task

Complete `Describe` in `describe.go` so `int` and `string` inputs are
named and everything else returns `"unknown"`.

## Check

```bash
go test ./exercises/04_interfaces/04_type_switch/ -v
```
