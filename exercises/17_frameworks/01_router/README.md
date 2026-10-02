# Routing with Gin

`net/http` ServeMux matches `"GET /tickets"`. Gin matches the same shape with `r.GET`:

```go
r := gin.New()
r.GET("/tickets", func(c *gin.Context) {
    c.String(200, "list")
})
```

`c.String` writes status + plain body. Unmatched paths 404 automatically — same default as ServeMux.

## What the docs add

Per [Gin routing](https://gin-gonic.com/en/docs/routing) and [HTTP methods](https://gin-gonic.com/en/docs/routing/http-method): routes are registered per HTTP verb — `GET`, `POST`, `PUT`, `PATCH`, `DELETE`, `HEAD`, `OPTIONS` — and each route answers **only** its own method. A `POST` to a `GET`-only path falls through to 404, exactly like ServeMux method patterns.

Two constructors, from the [quick start](https://gin-gonic.com/en/docs):

```go
r := gin.Default() // Logger + Recovery middleware attached
r := gin.New()      // bare engine, no middleware
```

These exercises use `gin.New()` plus `gin.SetMode(gin.TestMode)` in tests to keep output quiet. `gin.Default()` is what real servers reach for — its Recovery middleware is the framework version of panic-vs-error discipline.

<details>
<summary>Hint</summary>

One `r.GET("/tickets", ...)` writing `c.String(200, "list")`. Unmatched 404s need no code.

</details>

## Task

Fill in `NewRouter` in `router.go`:

```go
func NewRouter() *gin.Engine {
    // ...
}
```

Register `GET /tickets` to write `"list"`. The stub registers nothing.

## Check

```bash
go test ./exercises/17_frameworks/01_router/ -v
```
