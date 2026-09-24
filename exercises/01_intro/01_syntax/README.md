# Syntax

Welcome back — and no skipping! The warm-up in `exercises/01_intro/00_welcome`
should be green before you read on.

It hardly felt like an exercise, did it? Yet those few lines smuggled in
most of Go's **syntax**: packages, functions, parameters, return types.
Rather than dissecting every detail now, we'll unpack *just enough* to keep
you moving. The rest can wait until you actually need it.

## Comments

You can use `//` for single-line comments:

```go
// This is a single-line comment
// Followed by another single-line comment
```

Comments are for humans; the compiler ignores them. You'll see them marking
the exact line you need to change in each exercise's stub.

## Packages

Every Go file opens with a `package` clause:

```go
package syntax
```

Why must every file say it? Because Go compiles *packages*, not files. All
`.go` files in one folder must declare the same package name — disagree and
the build fails. Test files live in the same folder and share the name, which
is how they can call your functions directly.

## Functions

Functions are declared with the `func` keyword:

```go
// Syntax: func <name>(<parameters>) <return type> { <body> }
func Compute(a, b int) int {
    return a + b
}
```

`Compute` takes two inputs and produces one output, all of type `int`. The
name starts with a capital letter, which — as you'll learn in the ticket
section — makes it usable from other packages.

### Input parameters

Each parameter is a name plus a type. Adjacent parameters sharing a type can
share the annotation:

```go
// `a` and `b` are both `int`
func Compute(a, b int) int {
    return a + b
}
```

What if the types differ? Then each parameter carries its own:

```go
func Greet(name string, times int) string {
    if times <= 1 {
        return "Hello, " + name
    }
    return "Hello, " + name + "!"
}
```

Parameters are separated with commas. There is no limit, but if you find
yourself threading five strings through every call, that's a struct begging
to be born — you'll get there in the ticket section.

### Return type

The return type sits after the parameter list. What if a function returns
nothing at all? Then there is no annotation to write:

```go
func Log(msg string) {
    println(msg)
}
```

Go goes one further than most languages: a function can return *several*
values at once by wrapping them in parentheses. You'll meet this shape
constantly — it's how Go reports errors alongside results:

```go
func DivMod(a, b int) (int, int) {
    return a / b, a % b
}
```

Results can also be named, which pre-declares them as variables inside the
body:

```go
func DivModNamed(a, b int) (quot, rem int) {
    quot = a / b
    rem = a % b
    return quot, rem
}
```

### Returning values

Go returns values explicitly with `return` — there is no implicit "last
expression counts" rule. And the compiler holds you to the signature: every
path through a function that promises an `int` must `return` an `int`:

```go
func Compute(a, b int) int {
    // The value follows `return` on the same line.
    return a + b
}
```

Forget a path and the program doesn't build. Annoying today, beloved forever.

### Type annotations

Go is a **statically typed language**. Every value has a type, and the
compiler must know it when the program is built.

Think of a type as a **tag** the compiler pins to every value. The tag
decides the rules: you can add two numbers, but not a string to a number.
Lean on this — precise parameter and return types push entire families of
bugs from runtime (where users find them) to compile time (where you do).

One tag-related strictness to internalize now: inside a function, `:=`
declares a *new* variable while `=` assigns to an existing one, and a
declared-but-unused variable is a compile error, not a warning:

```go
sum := a + b // declared with `:=`, must be used below
```

## Task

Fix `Compute` in `syntax.go` so the test passes.

## Check

```bash
go test ./exercises/01_intro/01_syntax/ -v
```
