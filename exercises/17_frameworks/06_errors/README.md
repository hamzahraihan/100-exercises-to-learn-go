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

## From one handler to all of them

The `AbortWithStatusJSON` you write here handles one failure in one handler. But a ticket API has dozens of handlers, and each one can fail the same ways — not found, bad input, something broke. Copying the envelope into every handler rots the same way duplicated middleware does.

The way out is the post-`Next` phase from the middleware lesson, turned into a pattern. Handlers stop answering their own failures: they record them with `c.Error(err)` and keep going. One middleware, attached with `r.Use()`, runs after `c.Next()` returns, inspects the collected `c.Errors`, and writes the single envelope:

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

Same factoring-out as middleware itself, one level up: per-handler aborts become per-request error collection.

One related trap: `gin.CustomRecovery()` takes a `func(c *gin.Context, recovered any)` for panics, and inside it you must call an abort method yourself — `c.AbortWithStatus()`, for example. Forget it and subsequent handlers keep executing right past the panic.

## Further reading

- [Error-handling middleware](https://gin-gonic.com/en/docs/middleware/error-handling-middleware) in the Gin docs.

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
