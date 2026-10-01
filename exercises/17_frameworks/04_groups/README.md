# Groups with Gin

Real APIs version: `/api/v1/tickets`. Gin groups prefix in one call:

```go
v1 := r.Group("/api/v1")
v1.GET("/tickets", func(c *gin.Context) {
    c.String(200, "list")
})
```

Only grouped paths exist — bare `/tickets` 404s.

<details>
<summary>Hint</summary>

`r.Group("/api/v1")` then `v1.GET("/tickets", ...)` writing `"list"`.

</details>

## Task

Fill in `NewRouter` in `groups.go`.

## Check

```bash
go test ./exercises/17_frameworks/04_groups/ -v
```
