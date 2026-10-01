# Middleware with Gin

Stdlib wraps `http.Handler`. Gin runs `gin.HandlerFunc` in order, `c.Next()` delegating:

```go
func WithHeader() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Header("X-Course", "go")
        c.Next()
    }
}
r.Use(WithHeader())
```

Set-then-`Next`, same as set-then-delegate. Skip `Next` and the chain stops.

<details>
<summary>Hint</summary>

`c.Header("X-Course", "go")` before `c.Next()`.

</details>

## Task

Fill in `WithHeader` in `mw.go`.

## Check

```bash
go test ./exercises/17_frameworks/05_middleware/ -v
```
