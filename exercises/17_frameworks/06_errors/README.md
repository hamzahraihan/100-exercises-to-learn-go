# Errors with Gin

Stdlib writes `http.Error`. Gin aborts with JSON and stops the chain:

```go
r.GET("/tickets/:id", func(c *gin.Context) {
    if c.Param("id") != "7" {
        c.AbortWithStatusJSON(404, gin.H{"error": "not found"})
        return
    }
    c.JSON(200, gin.H{"id": 7})
})
```

The `return` after Abort is load-bearing — without it Gin double-writes.

<details>
<summary>Hint</summary>

Check `c.Param("id")`, `AbortWithStatusJSON(404, ...)` + `return`, else `c.JSON(200, ...)`.

</details>

## Task

Fill in `NewRouter` in `errors.go`.

## Check

```bash
go test ./exercises/17_frameworks/06_errors/ -v
```
