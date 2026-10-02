# Path Params with Gin

ServeMux uses `{id}` + `r.PathValue("id")`. Gin uses `:id` + `c.Param("id")`:

```go
r.GET("/tickets/:id", func(c *gin.Context) {
    c.String(200, "one:"+c.Param("id"))
})
```

The colon marks a segment hole. `c.Param` reads it back by name.

## Strict shapes

Reaching for `{id}` out of ServeMux habit? Nothing matches. Gin wants the colon form — `:name`, never `{name}` or `<name>` — and it is strict about shape: `/user/:name` matches `/user/john` but **not** `/user/` or `/user`. A missing segment isn't an empty value; it's a different route, and a different route means 404.

For catch-all tails there is the **wildcard** form, and it keeps the leading slash:

```go
router.GET("/user/:name/*action", func(c *gin.Context) {
    name := c.Param("name")     // "john"
    action := c.Param("action") // "/send"
})
```

## Further reading

- [Parameters in path](https://gin-gonic.com/en/docs/routing/param-in-path) in the Gin docs.

<details>
<summary>Hint</summary>

One `r.GET("/tickets/:id", ...)` writing `"one:"` plus `c.Param("id")`.

</details>

## Task

Fill in `NewRouter` in `params.go`:

```go
func NewRouter() *gin.Engine {
    // ...
}
```

## Check

```bash
go test ./exercises/17_frameworks/02_params/ -v
```
