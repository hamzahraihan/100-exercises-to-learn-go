# Syntax

Don't jump ahead!
Complete the exercise for the previous section before you start this one.
It's located in `exercises/01_intro/00_welcome`.

The previous task barely qualified as an exercise, but it already exposed you
to quite a bit of Go **syntax**. We won't cover every single detail used
there. Instead, we'll cover *just enough* to keep going without getting stuck.
One step at a time!

## Comments

You can use `//` for single-line comments:

```go
// This is a single-line comment
// Followed by another single-line comment
```

## Packages

Every Go file starts with a `package` clause. All `.go` files in one folder
must declare the same package name:

```go
package syntax
```

The package name is what other code uses to refer to your code. Test files
in the same folder share the package name so they can call your functions
directly.

## Functions

Functions are defined with the `func` keyword, followed by the function's
name, its input parameters, and its return type. The body is enclosed in
curly braces `{}`:

```go
// `func` <name> ( <input params> ) <return_type> { <body> }
func Compute(a, b int) int {
    return a + b
}
```

`Compute` takes two inputs and returns one output, all of type `int`.

### Input parameters

Each parameter is declared with its name followed by its type. Parameters
with the same type can share it:

```go
// `a` and `b` are both `int`
func Compute(a, b int) int {
    return a + b
}
```

If the types differ, each parameter needs its own type:

```go
//           👇            👇
func Greet(name string, times int) string {
    if times <= 1 {
        return "Hello, " + name
    }
    return "Hello, " + name + "!"
}
```

Multiple parameters are separated with commas.

### Return type

The return type comes after the parameter list. If the function returns
nothing, the return type is omitted entirely:

```go
func Log(msg string) {
    println(msg)
}
```

Functions can return more than one value by wrapping the results in
parentheses. You will see this constantly — it is how Go functions report
errors alongside their results:

```go
func DivMod(a, b int) (int, int) {
    return a / b, a % b
}
```

Results can also be named, which pre-declares them as variables:

```go
func DivModNamed(a, b int) (quot, rem int) {
    quot = a / b
    rem = a % b
    return quot, rem
}
```

### Returning values

Go returns values explicitly with the `return` keyword. Every code path in
a function with a return type must hit a `return` with a value of that
type — the compiler enforces this:

```go
func Compute(a, b int) int {
    // Notice: the value follows `return` on the same line.
    return a + b
}
```

### Type annotations

Go is a **statically typed language**. Every value has a type, and that type
must be known to the compiler when the program is built.

You can think of a type as a **tag** the compiler attaches to every value.
Depending on the tag, the compiler enforces different rules — you can't add
a string to a number, but you can add two numbers together. Used well, types
rule out whole classes of bugs before the program ever runs.

It is considered idiomatic to let types do this work for you: declare
precise parameter and return types rather than accepting a loose type and
checking at runtime.

## Common mistakes

- **Missing `return`:** a function declared to return `int` must return an
  `int` on every path. The compiler rejects the program otherwise.
- **Unused variables:** `x := 5` followed by never using `x` is a compile
  error. Remove the variable while experimenting.
- **`:=` vs `=`:** `:=` declares a *new* variable (`sum := a + b`), while
  `=` assigns to an existing one. `:=` only works inside functions.

## Task

Fix `Compute` in `syntax.go` so the test passes.

## Check

```bash
go test ./exercises/01_intro/01_syntax/ -v
```
