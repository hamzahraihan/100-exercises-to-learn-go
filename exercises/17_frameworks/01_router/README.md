# Routing with Gin

`net/http` ServeMux matches `"GET /tickets"`. Gin matches the same shape with `r.GET`:

```go
r := gin.New()
r.GET("/tickets", func(c *gin.Context) {
    c.String(200, "list")
})
```

`c.String` writes status + plain body. Unmatched paths 404 automatically — same default as ServeMux.

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
