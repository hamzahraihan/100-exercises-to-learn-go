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

## What the docs add

Per [model binding and validation](https://gin-gonic.com/en/docs/binding/binding-and-validation): Gin ships **two** binding families. `MustBind` methods (`Bind`, `BindJSON`) abort the request with 400 automatically on failure; `ShouldBind` methods (`ShouldBind`, `ShouldBindJSON`) return the error and let **you** decide the status and envelope — which is why this exercise uses the latter.

Validation can also move into the struct itself with `binding` tags ([FAQ](https://gin-gonic.com/en/docs/faq)):

```go
type User struct {
    Name  string `json:"name" binding:"required,min=3,max=50"`
    Email string `json:"email" binding:"required,email"`
    Age   int    `json:"age" binding:"gte=0,lte=130"`
}
```

`ShouldBindJSON` then enforces both shape *and* rules in one call. Two gotchas from the docs: fields must be **exported** (lowercase `age` never binds), and the returned `err.Error()` is safe to put in the 400 envelope since it describes the client's payload, not your internals.

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
