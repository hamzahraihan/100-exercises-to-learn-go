# Routing with ServeMux

`http.NewServeMux` routes requests to handlers by pattern. Since Go 1.22,
patterns can include the method (`"GET /tickets"`) and named path segments
(`"GET /tickets/{id}"`); inside the handler, `r.PathValue("id")` returns the
segment the client actually sent. A request that matches no pattern gets a
`404` automatically.

## Task

Fill in `NewMux` in `router.go`:

```go
func NewMux() *http.ServeMux {
	// ...
}
```

Register `"GET /tickets"` to write `"list"` and `"GET /tickets/{id}"` to
write `"one:"` plus the path value. The stub registers nothing, so the test
fails on the bodies.

## Check

```bash
go test ./exercises/13_http/04_routing/ -v
```
