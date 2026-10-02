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

## What the docs add

Per the [error-handling middleware](https://gin-gonic.com/en/docs/middleware/error-handling-middleware) docs: the per-handler `AbortWithStatusJSON` you write here scales into a centralized pattern. Handlers record failures with `c.Error(err)` and keep going; one middleware placed with `r.Use()` runs **after** `c.Next()` returns, inspects `c.Errors`, and writes a single envelope:

```go
func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next() // handlers run first
        if len(c.Errors) > 0 {
            c.JSON(500, gin.H{"success": false, "message": c.Errors.Last().Err.Error()})
        }
    }
}
```

That removes repetitive error handling from every handler — the same factoring-out that middleware itself performs. Related: `gin.CustomRecovery()` takes a `func(c *gin.Context, recovered any)` where you must call an abort method (e.g. `c.AbortWithStatus()`) or subsequent handlers keep executing past the panic.

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
