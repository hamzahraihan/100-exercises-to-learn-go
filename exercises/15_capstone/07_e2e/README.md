# End to End

Unit tests check one handler through a recorder; an end-to-end test boots the
real thing — `httptest.NewServer` binds one ephemeral socket, and the test
drives the full lifecycle over HTTP: `POST` a ticket (read its id from the
`201` reply), `GET` it, `PUT` an update, `DELETE` it, then `GET` again for
`404`. Here every handler and the store already work — the only broken layer
is the wiring: `NewServer` returns an empty mux, so everything `404`s. When a
single layer breaks, the e2e shape tells you exactly where to look.

## Task

Register the routes in `NewServer` in `e2e.go`:

```go
func NewServer() *Server {
	// ...
}
```

Wire `POST`+`GET /tickets` and `GET`+`PUT`+`DELETE /tickets/{id}` to the
existing handler methods, following the `s.mux.HandleFunc` pattern from the
earlier exercises. Nothing else needs to change.

## Check

```bash
go test ./exercises/15_capstone/07_e2e/ -v
```
