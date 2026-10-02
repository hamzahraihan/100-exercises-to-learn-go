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

## Three zoom levels

Logging belongs on every route. Auth belongs on a slice of them. A benchmark harness belongs on exactly one. Gin attaches middleware at all three zoom levels:

```go
router.Use(Logger(), Recovery())         // global: every route
v1.Use(AuthRequired())                   // group: everything under /v1
router.GET("/bench", BenchMw(), handler) // one route only
```

`gin.Default()` is nothing more than `gin.New()` with `Logger()` and `Recovery()` pre-attached at the global level.

## Before and after `Next`

`c.Next()` splits a middleware in two. Everything before it runs on the way in — headers, auth checks, setup. Everything after runs on the way out, with the response already written and observable. The canonical logger is the whole pattern in miniature:

```go
func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        c.Next() // handlers run here
        log.Printf("Request took %v", time.Since(start))
    }
}
```

Skip the `Next` call and the chain stops dead — requests answered mid-stack, handlers never reached. A middleware that never delegates isn't layered behavior; it's a wall.

## Further reading

- [Middleware](https://gin-gonic.com/en/docs/middleware) in the Gin docs.

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
