# Groups with Gin

Real APIs version: `/api/v1/tickets`. Gin groups prefix in one call:

```go
v1 := r.Group("/api/v1")
v1.GET("/tickets", func(c *gin.Context) {
    c.String(200, "list")
})
```

Only grouped paths exist — bare `/tickets` 404s.

## One route is a line; an API is a tree

A single versioned prefix fits in one call. A real API doesn't stay single for long — v1 beside v2, users beside posts, each with its own auth and logging. Without groups, every route repeats its full prefix and its middleware stack; the version string ends up copy-pasted a dozen times, and the day v2 arrives you get to rename it a dozen times.

Groups exist for exactly three jobs: a **shared prefix**, **shared middleware** for a whole slice of routes at once, and keeping related handlers visually together. And they nest:

```go
api := router.Group("/api")
v1 := api.Group("/v1")          // /api/v1/...
users := v1.Group("/users")     // /api/v1/users/...
users.GET("/", listUsers)
users.GET("/:id", getUser)
```

A middleware call on the group covers everything under it — `v1.Use(AuthRequired())` guards every v1 endpoint without touching a single handler. As the API grows, give each resource its own file, each registering on its own `gin.RouterGroup`: adding or removing a resource then can't disturb its neighbors.

## Further reading

- [Grouping routes](https://gin-gonic.com/en/docs/routing/grouping-routes) and [API design](https://gin-gonic.com/en/docs/routing/api-design) in the Gin docs.

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
