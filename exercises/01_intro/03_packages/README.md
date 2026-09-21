# Packages and imports

Go code lives in packages. `import "strings"` makes the standard library's
string helpers available, and only capitalized names (`ToUpper`) are visible
outside their package.

## Task

Fix `Shout` in `greet.go` so the test passes.

## Check

```bash
go test ./exercises/01_intro/03_packages/ -v
```
