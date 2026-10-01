# Path Params with Gin

ServeMux uses `{id}` + `r.PathValue("id")`. Gin uses `:id` + `c.Param("id")`:

```go
r.GET("/tickets/:id", func(c *gin.Context) {
    c.String(200, "one:"+c.Param("id"))
})
```

The colon marks a segment hole. `c.Param` reads it back by name.

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
