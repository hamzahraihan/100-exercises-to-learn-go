# Path Params with Gin

ServeMux uses `{id}` + `r.PathValue("id")`. Gin uses `:id` + `c.Param("id")`:

```go
r.GET("/tickets/:id", func(c *gin.Context) {
    c.String(200, "one:"+c.Param("id"))
})
```

The colon marks a segment hole. `c.Param` reads it back by name.

## What the docs add

Per [parameters in path](https://gin-gonic.com/en/docs/routing/param-in-path) and the [FAQ](https://gin-gonic.com/en/docs/faq): parameters use the colon prefix — `:name`, **not** `{name}` or `<name>`. Matching is strict about shape: `/user/:name` matches `/user/john` but **not** `/user/` or `/user`.

For catch-all tails there is the wildcard form:

```go
router.GET("/user/:name/*action", func(c *gin.Context) {
    name := c.Param("name")     // "john"
    action := c.Param("action") // "/send" — leading slash included
})
```

`/user/john/send` captures `action` as `"/send"`, leading slash and all.

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
