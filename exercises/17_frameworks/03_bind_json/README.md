# JSON Binding with Gin

Stdlib decodes with `json.NewDecoder(r.Body).Decode(&t)`. Gin binds in one call:

```go
var t Ticket
if err := c.ShouldBindJSON(&t); err != nil {
    c.JSON(400, gin.H{"error": "bad request"})
    return
}
if t.Title == "" {
    c.JSON(400, gin.H{"error": "title required"})
    return
}
c.JSON(201, t)
```

`ShouldBindJSON` (not `BindJSON`) returns the error for you to shape — same 400 envelope on malformed JSON and empty title.

<details>
<summary>Hint</summary>

`ShouldBindJSON`, check `Title`, `c.JSON(201, t)` on success else `c.JSON(400, ...)`.

</details>

## Task

Fill in `NewRouter` in `bind.go` with `POST /tickets` using the shape above.

## Check

```bash
go test ./exercises/17_frameworks/03_bind_json/ -v
```
