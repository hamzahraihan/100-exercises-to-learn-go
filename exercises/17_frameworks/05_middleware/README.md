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

## What the docs add

Per the [middleware docs](https://gin-gonic.com/en/docs/middleware): middleware attaches at three levels — globally, per group, or per route:

```go
router.Use(Logger(), Recovery()) // global: every route
v1.Use(AuthRequired())           // group: everything under /v1
router.GET("/bench", BenchMw(), handler) // one route only
```

`gin.Default()` is just `gin.New()` plus `Logger()` and `Recovery()` pre-attached. And `c.Next()` splits the function into pre- and post-phases — the canonical logger measures latency this way:

```go
func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        c.Next() // handlers run here
        log.Printf("Request took %v", time.Since(start))
    }
}
```

Code before `Next` is request setup (headers, auth checks); code after is response observation (status, latency, logging).

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
