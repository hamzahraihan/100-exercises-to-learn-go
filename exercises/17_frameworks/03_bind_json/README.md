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

## Who answers a failure?

Binding fails — malformed JSON, a wrong type, a missing field. Who decides what the client sees?

Gin ships **two** binding families, and they answer that question differently. The `MustBind` family (`Bind`, `BindJSON`) aborts the request with a 400 on failure, verdict included, no questions asked. The `ShouldBind` family (`ShouldBind`, `ShouldBindJSON`) hands the error back and lets **you** choose the status and the envelope. This exercise uses the latter, because a ticket API worthy of the name speaks in its own error shape, not the framework's default.

## Rules on the struct

Validation can move out of the handler entirely, into `binding` tags:

```go
type User struct {
    Name  string `json:"name" binding:"required,min=3,max=50"`
    Email string `json:"email" binding:"required,email"`
    Age   int    `json:"age" binding:"gte=0,lte=130"`
}
```

Now `ShouldBindJSON` enforces shape *and* rules in one call. Two traps, both silent. First: fields must be **exported** — a lowercase `age` never binds, and nothing complains. Second: the returned `err.Error()` is safe to echo in the 400 envelope, since it describes the client's payload, not your internals.

## Further reading

- [Model binding and validation](https://gin-gonic.com/en/docs/binding/binding-and-validation) in the Gin docs.

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
