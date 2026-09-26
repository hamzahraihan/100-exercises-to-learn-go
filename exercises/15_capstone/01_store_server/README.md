# Create and List

Everything so far arrives here: structs with tags, a mutex-guarded store,
handlers that decode and encode, status codes with meaning. This exercise
assembles them into the smallest thing deserving the name *API* — two
endpoints against one shared store.

## The band assembles

```go
type Server struct {
    store *Store
    mux   *http.ServeMux
}

func NewServer() *Server {
    s := &Server{store: &Store{}, mux: http.NewServeMux()}
    s.mux.HandleFunc("POST /tickets", s.handleCreate)
    s.mux.HandleFunc("GET /tickets", s.handleList)
    return s
}

func (s *Server) Handler() http.Handler { return s.mux }
```

Three roles, one struct. The **store** holds state (the mutex-guarded
slice list from the store sections, pre-written here — read it as the
specification for thread-safe habits: lock, touch, unlock, *copy on the
way out* via `List`). The **mux** routes (the routing lesson's patterns,
method included). The **handlers** are *methods* on `Server`, so
`s.store` is in scope wherever the logic needs it — receivers as
dependency injection, no globals, no parameters threaded through routing.

`Handler()` exposes the mux as an `http.Handler` for tests and, later,
real serving. Wiring lives in exactly one constructor; handlers know
nothing about routes, routes know nothing about logic.

## Two moves, full choreography

`handleCreate` performs the JSON lesson's triple — decode, judge, encode —
against the store:

```go
func (s *Server) handleCreate(w http.ResponseWriter, r *http.Request) {
    var t Ticket
    if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }
    if t.Title == "" {
        http.Error(w, "title required", http.StatusBadRequest)
        return
    }
    t = s.store.Add(t) // assigns the id, mutex-guarded
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(t)
}
```

Note what the handler *doesn't* do: no id generation (the store owns
identity), no locking (the store owns synchronization), no validation
beyond the border check. Handlers translate HTTP into store calls and
store results into HTTP — thin by design. `handleList` is thinner still:
encode `s.store.List()` with `200`, three lines where the only decision
is the content type.

The stubs answer `501 Not Implemented` — the honest placeholder for
"route exists, logic doesn't." Tests drive both endpoints through the
mux (`ServeHTTP` directly, recorder capturing), proving wiring *and*
logic in one pass: POST returns 201, GET lists the created ticket.

## Task

Fill in `handleCreate` and `handleList` in `server.go`:

```go
func (s *Server) handleCreate(w http.ResponseWriter, r *http.Request) {
    // ...
}

func (s *Server) handleList(w http.ResponseWriter, r *http.Request) {
    // ...
}
```

Decode with `json.NewDecoder(r.Body).Decode(&v)`, reject an empty title with
`400`, otherwise `store.Add` and reply `201` with the stored ticket as JSON.
For the list, encode `s.store.List()` as JSON with `200`. Set
`w.Header().Set("Content-Type", "application/json")` and call `WriteHeader`
before encoding. The stubs answer `501`, so the tests fail on status.

## Check

```bash
go test ./exercises/15_capstone/01_store_server/ -v
```
