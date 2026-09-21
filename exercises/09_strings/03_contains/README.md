# Substring Checks

Often you only need to ask where a string sits inside another: does it contain
this piece, start with that prefix, end with that suffix? The `strings`
package names one function for each — `Contains`, `HasPrefix`, `HasSuffix` —
so the call reads like the question. The check here is deliberately naive (a
real validator would do much more), and the stub always says `false`, so the
table test fails on the positive rows.

## Task

Fill in `IsEmail` in `mail.go`:

```go
func IsEmail(s string) bool {
	// ...
}
```

Use `strings.Contains` to report whether the address holds an `"@"`.

## Check

```bash
go test ./exercises/09_strings/03_contains/ -v
```
