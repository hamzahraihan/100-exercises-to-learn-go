# Packages and imports

So far every exercise lived in a single file. Real programs don't — and Go
has opinions about how files find each other. This exercise's stub is one
line, but it opens two doors: the standard library and the visibility rule
that governs all of Go.

## Borrowing from the standard library

`greet.go` starts with a line you haven't written yourself yet:

```go
import "strings"
```

That string is an **import path**, and it buys you the entire `strings`
package: case conversion, trimming, splitting, joining, searching. Go ships
with a standard library most languages would call a framework — and using
it is always one import away, no package manager involved.

The exercise asks for shouting, and the library already knows how:

```go
// Syntax: <package>.<Name>(<args>)
strings.ToUpper("hello") // "HELLO"
strings.ToLower("HELLO") // "hello"
```

Notice the stub's bug with fresh eyes: it calls `ToLower` where the test
demands upper case. The fix is a single word. Most real bugs you meet in
this course will be exactly this shape — the right library, the wrong
member — which is why reading compiler and test output precisely beats
guessing.

## The capital-letter rule

Look closely at those names: `ToUpper`, not `toUpper`. In Go, **a name
starting with a capital letter is exported** — visible to any package that
imports it. Lowercase names stay private to their own package.

```go
strings.ToUpper // ✅ exported: you may call this
strings.toUpper // ❌ doesn't exist outside package strings
```

There are no `public` or `private` keywords anywhere in the language. The
first letter *is* the access modifier. It applies uniformly — functions,
types, struct fields, methods:

- `Shout` can be called by the test in the same breath as by strangers.
- `Ticket.Title` (capital T) will be readable across packages when tickets
  arrive; a lowercase field would not be.

One rule, no exceptions, enforced by the compiler. When an "undefined"
error names something you can *see* right there in the source, check the
casing first — you're likely knocking on a private door.

## What `import` really does

A few mechanics worth knowing now, before the HTTP section leans on them:

- **Unused imports don't compile.** Import `strings` and never call it,
  and the build fails. The compiler would rather nag than ship dead weight.
- **One import per package.** Importing `"strings"` twice, or importing
  something the file never uses "just in case," is rejected outright.
- **The package name can differ from the path's last word**, though it
  usually matches. When it doesn't, you call it by its declared name, not
  the path. You'll see this with versioned paths later.

Grouped imports in parentheses are the norm once a file needs several:

```go
import (
    "strings"
    "testing"
)
```

The test file has been doing this quietly all along — `testing` is itself
just another package, borrowed the same way.

## Task

Fix `Shout` in `greet.go` so the test passes.

## Check

```bash
go test ./exercises/01_intro/03_packages/ -v
```
