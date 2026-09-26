# End to End

Unit tests proved each handler through a recorder. But recorders don't
route, serialize across sockets, or read real status lines — the seams
*between* layers go untested. This finale boots the actual server over an
actual (local) socket and drives the full ticket lifecycle: create it,
read it, change it, delete it, prove it's gone.

## The whole band, live

```go
srv := httptest.NewServer(NewServer().Handler())
defer srv.Close()
```

`httptest.NewServer` binds an ephemeral port and serves for real — actual
HTTP over localhost, unlike the recorder's direct function calls. Nothing
leaves the machine, but everything else is genuine: TCP, headers on the
wire, status lines parsed back. `defer srv.Close()` tears it down; the
server lives exactly as long as the test.

The test then walks the lifecycle, each step's output feeding the next:

```go
POST   /tickets     {"title":"A",...} → 201, reply body carries the id
GET    /tickets/{id}                  → 200, body contains "title":"A"
PUT    /tickets/{id}  {"title":"B",...} → 200
DELETE /tickets/{id}                  → 204
GET    /tickets/{id}                  → 404, the after-delete proof
```

Create-then-read proves persistence across requests (the store survives
beyond one handler call). The created id is *parsed from the reply*
(`json.Unmarshal` into an id struct) rather than assumed — the test
discovers state the way a client would. Update-then-delete-then-404 closes
the loop the delete lesson opened: absence verified through the front
door, over the wire.

## The driver helper

```go
func do(t *testing.T, method, url, body string) (int, string) {
    t.Helper()
    // ... build request, run it with http.DefaultClient.Do ...
    defer resp.Body.Close()
    raw, _ := io.ReadAll(resp.Body)
    return resp.StatusCode, string(raw)
}
```

One helper for five calls: method, URL, optional body in; status code and
body out. `t.Helper()` marks the infrastructure (the helpers lesson,
unchanged), the real client `Do` sends (the client lesson, unchanged),
and the body closes via defer (the close-first sacrament, unchanged). The
e2e test is a composition of course habits as much as of handlers — every
line in it was taught somewhere earlier.

## When one layer breaks

Here everything works except the wiring: `NewServer` returns an empty mux,
so every request 404s. That isolation is deliberate — and diagnostic. When
this suite fails, the failure *location* names the layer: all-404s means
routing, a 400 on PUT means validation, a missing id means creation. Your
task is five `HandleFunc` lines following the established pattern
(`POST`+`GET /tickets`, `GET`+`PUT`+`DELETE /tickets/{id}`), and nothing
else changes. The last exercise of the course asks for the least new
knowledge — only the confidence that the pieces fit, because you watched
each one being built.

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
