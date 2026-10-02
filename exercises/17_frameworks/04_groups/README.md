# Groups with Gin

Real APIs version: `/api/v1/tickets`. Gin groups prefix in one call:

```go
v1 := r.Group("/api/v1")
v1.GET("/tickets", func(c *gin.Context) {
    c.String(200, "list")
})
```

Only grouped paths exist — bare `/tickets` 404s.

## What the docs add

Per [grouping routes](https://gin-gonic.com/en/docs/routing/grouping-routes) and [API design](https://gin-gonic.com/en/docs/routing/api-design): groups exist for three jobs — shared prefixes (versioning as `/api/v1`, `/api/v2` side by side), shared middleware for a whole slice of routes at once, and keeping related handlers visually grouped. Groups nest:

```go
api := router.Group("/api")
v1 := api.Group("/v1")          // /api/v1/...
users := v1.Group("/users")     // /api/v1/users/...
users.GET("/", listUsers)
users.GET("/:id", getUser)
```

Middleware can attach at group level (`v1.Use(AuthRequired())`), so auth/versioning/logging apply to everything under the prefix without repeating the call. As APIs grow, the docs recommend one file per resource, each registering on its own `gin.RouterGroup`.

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
