# Routing with Gin

`net/http` ServeMux matches `"GET /tickets"`. Gin matches the same shape with `r.GET`:

```go
r := gin.New()
r.GET("/tickets", func(c *gin.Context) {
    c.String(200, "list")
})
```

`c.String` writes status + plain body. Unmatched paths 404 automatically — same default as ServeMux.

## One method per route

A Gin route answers exactly one HTTP verb. Register `GET /tickets` and a `POST` to that same path falls through to 404 — the method is part of the match, not decoration. The full family is there when you need it: `GET` for reading, `POST` for creating, `PUT` for replacing, `PATCH` for partial updates, `DELETE` for removing, plus `HEAD` and `OPTIONS`.

## Two constructors

```go
r := gin.Default() // Logger + Recovery middleware attached
r := gin.New()      // bare engine, no middleware
```

`gin.Default()` is what real servers reach for — its Recovery middleware is the framework's answer to the panic-vs-error discipline. These exercises use `gin.New()`, with the tests silencing debug output via `gin.SetMode(gin.TestMode)`.

## Further reading

- [Routing](https://gin-gonic.com/en/docs/routing) and [HTTP methods](https://gin-gonic.com/en/docs/routing/http-method) in the Gin docs.

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
