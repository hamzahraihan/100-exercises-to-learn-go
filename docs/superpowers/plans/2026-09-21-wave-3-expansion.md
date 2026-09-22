# Wave 3 Implementation Plan (78 → 100 exercises)

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add the final 22 exercises (13_http, 14_persistence, 15_capstone ticket API server), extending `solutions` to 100/100.

**Architecture:** Same course architecture: single module `learngo`, per-exercise triplet (broken-by-design stub compiling under `go vet` + failing test + teaching README). HTTP tests use `httptest` in-memory (only the e2e binds a socket via `httptest.NewServer`); persistence uses `t.TempDir()`; capstone dirs carry working store scaffolding with broken handlers as the exercise. Sections never import each other.

**Tech Stack:** Go >= 1.23 (1.22+ ServeMux patterns + PathValue required), stdlib only (`net/http`, `net/http/httptest`, `encoding/json`, `os`, `context`, `time`, `strconv`), unchanged CI.

## Global Constraints

- Go floor: `go 1.23` in `go.mod`; module path exactly `learngo`.
- Test deps: stdlib `testing` (+ `httptest`) only; no external dependencies, no third-party routers.
- Exercise contract: stub `.go` MUST compile under `go vet ./...` (all imports used; unused function params are fine); its test MUST fail on assertion/status/body, never on build.
- READMEs teach without printing full solutions: signatures + `// ...` only; naming stdlib items is allowed.
- Determinism: `httptest` recorders/servers (ephemeral ports), `t.TempDir()`, fixed fixtures; assert outcomes never durations; generous margins (>= 10x).
- `*.go` + fixtures LF via `.gitattributes`; `go vet`+`go build` green on `main`, `go test` red on `main` by design, green on `solutions`.
- All prose original Go teaching; attribution already in root README.

---

### Task 1: New section `13_http` (+10 → 88)

**Files:** New dirs `01_handler` (hello.go, hello_test.go), `02_status` (create.go, create_test.go), `03_query` (filter.go, filter_test.go), `04_routing` (router.go, router_test.go), `05_json_api` (api.go, api_test.go), `06_httptest` (ping.go, ping_test.go), `07_middleware` (mw.go, mw_test.go), `08_context` (greet.go, greet_test.go), `09_client` (getbody.go, getbody_test.go), `10_error_response` (errw.go, errw_test.go), each + `README.md`. Packages: `hello`, `created`, `filter`, `router`, `api`, `ping`, `mw`, `greetctx`, `getbody`, `errw`.

**Interfaces:**
- Consumes: nothing.
- Produces: `func Hello(w, r)`, `func Create(w, r)` (201), `func StatusParam(r) string`, `func NewMux() *http.ServeMux`, Ticket JSON `func CreateTicket(w, r)`, `func Ping(w, r)`, `func WithHeader(next http.Handler) http.Handler`, `func Greet(ctx, name) (string, error)`, `func GetBody(url string) (string, error)`, `func WriteError(w, code, msg)`.

- [ ] **Step 1: `01_handler` + `02_status`**

`hello_test.go`:
```go
package hello

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	rec := httptest.NewRecorder()
	Hello(rec, httptest.NewRequest(http.MethodGet, "/", nil))
	if rec.Body.String() != "hello" {
		t.Fatalf("body = %q, want %q", rec.Body.String(), "hello")
	}
}
```

`hello.go`:
```go
// Package hello teaches http.HandlerFunc basics.
package hello

import "net/http"

// Hello writes "hello". Unused params are fine; the imports must be used.
// TODO: write the body (hint: Fprintf from fmt, or io.WriteString).
func Hello(w http.ResponseWriter, r *http.Request) {
}
```

`create_test.go`:
```go
package created

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	Create(rec, httptest.NewRequest(http.MethodPost, "/tickets", nil))
	if rec.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusCreated)
	}
}
```

`create.go`:
```go
// Package created teaches status codes.
package created

import "net/http"

// Create answers resource creation.
// TODO: w.WriteHeader(http.StatusCreated).
func Create(w http.ResponseWriter, r *http.Request) {
}
```
READMEs: HandlerFunc signature, recorder/request pattern; status codes + WriteHeader-before-write rule.

- [ ] **Step 2: `03_query` + `04_routing`**

`filter_test.go`:
```go
package filter

import (
	"net/http/httptest"
	"testing"
)

func TestStatusParam(t *testing.T) {
	r := httptest.NewRequest("GET", "/tickets?status=open", nil)
	if got, want := StatusParam(r), "open"; got != want {
		t.Fatalf("StatusParam = %q, want %q", got, want)
	}
	r = httptest.NewRequest("GET", "/tickets", nil)
	if got, want := StatusParam(r), "all"; got != want {
		t.Fatalf("StatusParam = %q, want %q", got, want)
	}
}
```

`filter.go`:
```go
// Package filter teaches query parameters.
package filter

import "net/http"

// StatusParam returns the "status" query value, defaulting to "all".
// TODO: r.URL.Query().Get with an empty check.
func StatusParam(r *http.Request) string {
	return ""
}
```

`router_test.go`:
```go
package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func get(t *testing.T, mux http.Handler, path string) *httptest.ResponseRecorder {
	t.Helper()
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
	return rec
}

func TestRouting(t *testing.T) {
	mux := NewMux()
	if rec := get(t, mux, "/tickets"); rec.Body.String() != "list" {
		t.Fatalf("GET /tickets = %q, want list", rec.Body.String())
	}
	if rec := get(t, mux, "/tickets/7"); rec.Body.String() != "one:7" {
		t.Fatalf("GET /tickets/7 = %q, want one:7", rec.Body.String())
	}
}
```

`router.go` (imports only what the stub uses — the fix may add `fmt`):
```go
// Package router teaches ServeMux method+path patterns (Go 1.22+).
package router

import "net/http"

// NewMux routes the ticket collection and single-ticket reads.
// TODO: mux.HandleFunc("GET /tickets", ...) writing "list" and
// "GET /tickets/{id}" writing "one:"+r.PathValue("id").
func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
```
READMEs: query parsing; 1.22 patterns + PathValue (describe).

- [ ] **Step 3: `05_json_api` + `06_httptest`**

`api_test.go`:
```go
package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateTicket(t *testing.T) {
	rec := httptest.NewRecorder()
	body := strings.NewReader(`{"id":1,"title":"Fix bug","status":"open"}`)
	CreateTicket(rec, httptest.NewRequest(http.MethodPost, "/tickets", body))
	if rec.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusCreated)
	}
	if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
		t.Fatalf("Content-Type = %q, want application/json", ct)
	}
	if !strings.Contains(rec.Body.String(), `"title":"Fix bug"`) {
		t.Fatalf("body = %s, want title echoed", rec.Body.String())
	}
}

func TestCreateTicketBad(t *testing.T) {
	rec := httptest.NewRecorder()
	body := strings.NewReader(`{"id":1,"status":"open"}`)
	CreateTicket(rec, httptest.NewRequest(http.MethodPost, "/tickets", body))
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
}
```

`api.go` (Ticket shape setup + empty stub; no unused imports — stub imports only `net/http`):
```go
// Package api teaches JSON request/response handling.
package api

import "net/http"

// Ticket is a support request.
type Ticket struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status"`
}

// CreateTicket decodes a ticket, validates the title, and echoes it back
// with 201 as JSON (400 for bad input).
// TODO: json.Decode the body, check Title, set Content-Type, WriteHeader,
// json.Encode the ticket (import encoding/json).
func CreateTicket(w http.ResponseWriter, r *http.Request) {
}
```
README: decode/encode, Content-Type, status-before-body rule.

`ping_test.go` (the test file itself demonstrates the httptest pattern):
```go
package ping

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	rec := httptest.NewRecorder()
	Ping(rec, httptest.NewRequest(http.MethodGet, "/ping", nil))
	if rec.Body.String() != "pong" {
		t.Fatalf("body = %q, want pong", rec.Body.String())
	}
}
```
Stub `ping.go`: empty `Ping` (fails). README: NewRecorder/NewRequest anatomy, table of recorder inspectors (Code/Header/Body).

- [ ] **Step 4: `07_middleware` + `08_context`**

`mw_test.go`:
```go
package mw

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWithHeader(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	rec := httptest.NewRecorder()
	WithHeader(next).ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))
	if rec.Header().Get("X-Course") != "go" {
		t.Fatalf("X-Course = %q, want go", rec.Header().Get("X-Course"))
	}
	if rec.Body.String() != "ok" {
		t.Fatalf("body = %q, want ok", rec.Body.String())
	}
}
```

`mw.go`:
```go
// Package mw teaches middleware.
package mw

import "net/http"

// WithHeader tags every response with X-Course: go, then calls next.
// TODO: w.Header().Set(...) before next.ServeHTTP(w, r).
func WithHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
```
README: `func(http.Handler) http.Handler` shape, chaining order.

`greet_test.go`:
```go
package greetctx

import (
	"context"
	"testing"
)

func TestGreet(t *testing.T) {
	got, err := Greet(context.Background(), "bob")
	if err != nil || got != "hi bob" {
		t.Fatalf("Greet = (%q, %v), want (hi bob, nil)", got, err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := Greet(ctx, "bob"); err == nil {
		t.Fatal("Greet(canceled) = nil error, want cancellation")
	}
}
```

`greet.go`:
```go
// Package greetctx teaches request contexts.
package greetctx

import (
	"context"
	"time"
)

// Greet greets after a 10ms simulated lookup, aborting on cancellation.
// TODO: select on ctx.Done() (return ctx.Err()) vs time.After result.
func Greet(ctx context.Context, name string) (string, error) {
	<-time.After(10 * time.Millisecond)
	return "hi " + name, nil
}
```
README: contexts carry cancellation, always select on Done (describe).

- [ ] **Step 5: `09_client` + `10_error_response`**

`getbody_test.go`:
```go
package getbody

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("data"))
	}))
	defer srv.Close()
	got, err := GetBody(srv.URL)
	if err != nil || got != "data" {
		t.Fatalf("GetBody = (%q, %v), want (data, nil)", got, err)
	}
}
```

`getbody.go`:
```go
// Package getbody teaches outgoing requests.
package getbody

// GetBody fetches url and returns its body as a string.
// TODO: http.Get, check err, defer Body.Close, io.ReadAll (import net/http, io).
func GetBody(url string) (string, error) {
	return "", nil
}
```
README: close bodies, `NewRequestWithContext` mentioned (no sleeps; local server only).

`errw_test.go`:
```go
package errw

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWriteError(t *testing.T) {
	rec := httptest.NewRecorder()
	WriteError(rec, http.StatusNotFound, "nope")
	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
		t.Fatalf("Content-Type = %q, want application/json", ct)
	}
	if !strings.Contains(rec.Body.String(), "nope") {
		t.Fatalf("body = %s, want message inside", rec.Body.String())
	}
}
```

`errw.go`:
```go
// Package errw teaches structured error responses.
package errw

import "net/http"

// WriteError answers code with a {"error":msg} JSON body.
// TODO: header, WriteHeader(code), encode map/struct (import encoding/json).
func WriteError(w http.ResponseWriter, code int, msg string) {
}
```
README: error envelope shape, status-before-body (describe).

- [ ] **Step 6: Verify and commit**

Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test ./exercises/13_http/...`
Expected: FAIL (10 packages on status/body assertions; no sockets bound except `09_client`'s ephemeral `httptest.NewServer`).

```bash
git add exercises/13_http
git commit -m "feat: add 13_http section (wave 3)"
```

---

### Task 2: New section `14_persistence` (+5 → 93)

**Files:** New dirs `01_save` (atomic.go, atomic_test.go), `02_load` (load.go, load_test.go), `03_corrupt` (loadtickets.go, loadtickets_test.go), `04_export` (export.go, export_test.go), `05_import` (importtk.go, importtk_test.go), each + `README.md`. Packages: `atomic`, `loadst`, `loadj`, `exporttk`, `importtk`. Ticket shape where needed: `{ID int, Title string}` (+ JSON tags in 04/05).

**Interfaces:**
- Consumes: nothing.
- Produces: `func Save(path string, data []byte) error`, `func Load(path string) ([]byte, error)`, `func LoadTickets(path string) ([]Ticket, error)`, `func Export(path string, ts []Ticket) error`, `func Import(path string) ([]Ticket, error)`.

- [ ] **Step 1: `01_save` + `02_load`**

`atomic_test.go`:
```go
package atomic

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSave(t *testing.T) {
	path := filepath.Join(t.TempDir(), "data.json")
	if err := Save(path, []byte(`{"id":1}`)); err != nil {
		t.Fatalf("Save errored: %v", err)
	}
	got, err := os.ReadFile(path)
	if err != nil || string(got) != `{"id":1}` {
		t.Fatalf("file = (%q, %v)", got, err)
	}
}
```

`atomic.go`:
```go
// Package atomic teaches crash-safe file writes.
package atomic

// Save writes data atomically: temp file in the same dir + rename, so a
// crash never leaves a half-written file.
// TODO: os.CreateTemp(dir...), Write, Close, os.Rename (import os, path/filepath).
func Save(path string, data []byte) error {
	return nil
}
```
README: why rename is atomic, same-dir temp requirement (describe).

`load_test.go`:
```go
package loadst

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "data.json")
	if err := os.WriteFile(path, []byte("hi"), 0o644); err != nil {
		t.Fatal(err)
	}
	if got, err := Load(path); err != nil || string(got) != "hi" {
		t.Fatalf("Load = (%q, %v)", got, err)
	}
	if got, err := Load(filepath.Join(dir, "missing.json")); err != nil || got != nil {
		t.Fatalf("Load(missing) = (%q, %v), want (nil, nil)", got, err)
	}
}
```

`load.go`:
```go
// Package loadst teaches boot-time loading.
package loadst

import "os"

// Load reads path; a missing file means empty state, not an error.
// TODO: os.ReadFile, map os.IsNotExist(err) to (nil, nil)
// (import errors only if needed for errors.Is style).
func Load(path string) ([]byte, error) {
	return os.ReadFile(path)
}
```
README: missing-file-as-empty convention, `os.IsNotExist` (describe).

- [ ] **Step 2: `03_corrupt` + `04_export` + `05_import`**

`loadtickets_test.go`:
```go
package loadj

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadTickets(t *testing.T) {
	dir := t.TempDir()
	ok := filepath.Join(dir, "ok.json")
	os.WriteFile(ok, []byte(`[{"id":1,"title":"a"}]`), 0o644)
	got, err := LoadTickets(ok)
	if err != nil || len(got) != 1 || got[0].Title != "a" {
		t.Fatalf("LoadTickets = (%+v, %v)", got, err)
	}
	bad := filepath.Join(dir, "bad.json")
	os.WriteFile(bad, []byte(`{oops`), 0o644)
	if _, err := LoadTickets(bad); err == nil {
		t.Fatal("corrupt file accepted, want descriptive error")
	} else if !strings.Contains(err.Error(), bad) {
		t.Fatalf("error = %q, want path inside", err.Error())
	}
}
```

`loadtickets.go` (Ticket setup + stub returning nil,nil):
```go
// Package loadj teaches graceful corrupt-file handling.
package loadj

// Ticket is a stored support request.
type Ticket struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// LoadTickets decodes path; failures name the path (never panic).
// TODO: os.ReadFile + json.Unmarshal, wrap errors with path (import encoding/json, fmt, os).
func LoadTickets(path string) ([]Ticket, error) {
	return nil, nil
}
```
README: wrap with context (`%w` + path), never panic on input.

`export_test.go`:
```go
package exporttk

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestExport(t *testing.T) {
	path := filepath.Join(t.TempDir(), "backup.json")
	in := []Ticket{{ID: 1, Title: "a"}, {ID: 2, Title: "b"}}
	if err := Export(path, in); err != nil {
		t.Fatalf("Export errored: %v", err)
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("backup missing: %v", err)
	}
	var got []Ticket
	if err := json.Unmarshal(raw, &got); err != nil || len(got) != 2 || got[1].Title != "b" {
		t.Fatalf("backup = (%s, %v)", raw, err)
	}
}
```

Stub (`export.go`, Ticket `{ID int \`json:"id"\`, Title string \`json:"title"\`}`):
```go
// Export writes ts as a JSON array to path.
// TODO: json.Marshal + os.WriteFile(path, out, 0o644) (import encoding/json, os).
func Export(path string, ts []Ticket) error {
	return nil
}
```
README: export format = JSON array, pretty vs compact (describe).

`importtk_test.go`:
```go
package importtk

import (
	"os"
	"path/filepath"
	"testing"
)

func TestImport(t *testing.T) {
	dir := t.TempDir()
	ok := filepath.Join(dir, "ok.json")
	os.WriteFile(ok, []byte(`[{"id":1,"title":"a"}]`), 0o644)
	got, err := Import(ok)
	if err != nil || len(got) != 1 || got[0].Title != "a" {
		t.Fatalf("Import = (%+v, %v)", got, err)
	}
	bad := filepath.Join(dir, "bad.json")
	os.WriteFile(bad, []byte(`[{"id":1,"title":""}]`), 0o644)
	if _, err := Import(bad); err == nil {
		t.Fatal("empty title imported, want validation error")
	}
}
```

Stub (`importtk.go`, same Ticket shape):
```go
// Import reads a JSON backup, validating every title.
// TODO: os.ReadFile + json.Unmarshal, reject empty titles (import encoding/json, errors, os).
func Import(path string) ([]Ticket, error) {
	return nil, nil
}
```
README: validate on import, replace-vs-merge choice (describe: replace).

- [ ] **Step 3: Verify and commit**

Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test ./exercises/14_persistence/...`
Expected: FAIL (5 packages; missing-file/corrupt cases fail by design).
Confirm: `git status --short` shows only the 15 intended files (TempDir leaves nothing).

```bash
git add exercises/14_persistence
git commit -m "feat: add 14_persistence section (wave 3)"
```

---

### Task 3: New section `15_capstone` (+7 → 100)

**Files:** New dirs `01_store_server` (server.go, server_test.go), `02_get_one` (getone.go, getone_test.go), `03_update` (update.go, update_test.go), `04_delete` (delete.go, delete_test.go), `05_graceful` (shutdown.go, shutdown_test.go), `06_config` (config.go, config_test.go), `07_e2e` (e2e.go, e2e_test.go), each + `README.md`. Packages: `server`, `getone`, `update`, `del`, `shutdown`, `config`, `e2e`. Each dir self-contained: working Store scaffolding + broken handler(s) as the exercise (01–04, 07); standalone funcs for 05–06.

**Interfaces:**
- Consumes: nothing (Ticket/Store shapes redeclared per dir — same JSON shape as Wave 2).
- Produces: wired `Server` with CRUD handlers, `ShutdownGracefully`, `ConfigFromEnv`, full e2e flow.

Shared scaffolding used VERBATIM in 01–04 and 07 (working code in stub, NOT the exercise — copy exactly; only the noted handler/route TODOs differ per dir; each dir imports `sync` for it):
```go
type Ticket struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status"`
}

type Store struct {
	mu    sync.Mutex
	items []Ticket
	next  int
}

func (s *Store) Add(t Ticket) Ticket {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.next++
	t.ID = s.next
	s.items = append(s.items, t)
	return t
}

func (s *Store) Get(id int) (Ticket, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, t := range s.items {
		if t.ID == id {
			return t, true
		}
	}
	return Ticket{}, false
}

func (s *Store) Update(id int, t Ticket) (Ticket, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.items {
		if s.items[i].ID == id {
			t.ID = id
			s.items[i] = t
			return t, true
		}
	}
	return Ticket{}, false
}

func (s *Store) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.items {
		if s.items[i].ID == id {
			s.items = append(s.items[:i], s.items[i+1:]...)
			return true
		}
	}
	return false
}
```

- [ ] **Step 1: `01_store_server` (create + list)**

`server_test.go`:
```go
package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateAndList(t *testing.T) {
	s := NewServer()
	rec := httptest.NewRecorder()
	body := strings.NewReader(`{"title":"Fix bug","status":"open"}`)
	s.Handler().ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/tickets", body))
	if rec.Code != http.StatusCreated {
		t.Fatalf("POST status = %d, want 201", rec.Code)
	}
	rec = httptest.NewRecorder()
	s.Handler().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/tickets", nil))
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), "Fix bug") {
		t.Fatalf("GET = (%d, %s)", rec.Code, rec.Body.String())
	}
}
```

`server.go`: working `Ticket` + `Store{Add, List}` + `Server{store, mux}` + `NewServer` wiring `POST /tickets` → handleCreate, `GET /tickets` → handleList + `Handler()` accessor. Broken handlers return 501:
```go
func (s *Server) handleCreate(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "TODO", http.StatusNotImplemented)
}

func (s *Server) handleList(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "TODO", http.StatusNotImplemented)
}
```
README: wiring store to handlers, 201 + JSON echo (describe create validation: non-empty title → 400).

- [ ] **Step 2: `02_get_one` + `03_update` + `04_delete`**

- [ ] **Step 2: `02_get_one` + `03_update` + `04_delete`**

Same scaffolding pattern per dir (copy the shared `Ticket`+`Store` verbatim, plus
`type Server struct{ store *Store; mux *http.ServeMux }`, a `NewServer` wiring
that dir's routes, and `func (s *Server) Handler() http.Handler { return s.mux }`;
only the target handler returns 501 as the exercise). Stub handler shape (example):
```go
func (s *Server) handleGet(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "TODO", http.StatusNotImplemented)
}
```

`getone_test.go` (test seeds via direct `store.Add`; `Add` numbers from 1, so the
first ticket is id 1):
```go
func TestGetOne(t *testing.T) {
	s := NewServer()
	s.store.Add(Ticket{Title: "A", Status: "open"})
	rec := httptest.NewRecorder()
	s.Handler().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/tickets/1", nil))
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), "A") {
		t.Fatalf("GET /tickets/1 = (%d, %s)", rec.Code, rec.Body.String())
	}
	rec = httptest.NewRecorder()
	s.Handler().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/tickets/99", nil))
	if rec.Code != http.StatusNotFound {
		t.Fatalf("GET /tickets/99 = %d, want 404", rec.Code)
	}
}
```
(Use this version — no unused variables.)

`update_test.go`:
```go
func TestUpdate(t *testing.T) {
	s := NewServer()
	s.store.Add(Ticket{Title: "A", Status: "open"})
	rec := httptest.NewRecorder()
	body := strings.NewReader(`{"title":"B","status":"closed"}`)
	s.Handler().ServeHTTP(rec, httptest.NewRequest(http.MethodPut, "/tickets/1", body))
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), `"title":"B"`) {
		t.Fatalf("PUT = (%d, %s)", rec.Code, rec.Body.String())
	}
	rec = httptest.NewRecorder()
	s.Handler().ServeHTTP(rec, httptest.NewRequest(http.MethodPut, "/tickets/99", strings.NewReader(`{}`)))
	if rec.Code != http.StatusNotFound {
		t.Fatalf("PUT missing = %d, want 404", rec.Code)
	}
	rec = httptest.NewRecorder()
	s.Handler().ServeHTTP(rec, httptest.NewRequest(http.MethodPut, "/tickets/1", strings.NewReader(`{"title":""}`)))
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("PUT empty title = %d, want 400", rec.Code)
	}
}
```
(imports `net/http`, `net/http/httptest`, `strings`, `testing`; package `update`.)

`delete_test.go`:
```go
func TestDelete(t *testing.T) {
	s := NewServer()
	s.store.Add(Ticket{Title: "A", Status: "open"})
	rec := httptest.NewRecorder()
	s.Handler().ServeHTTP(rec, httptest.NewRequest(http.MethodDelete, "/tickets/1", nil))
	if rec.Code != http.StatusNoContent {
		t.Fatalf("DELETE = %d, want 204", rec.Code)
	}
	rec = httptest.NewRecorder()
	s.Handler().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/tickets/1", nil))
	if rec.Code != http.StatusNotFound {
		t.Fatalf("GET after delete = %d, want 404", rec.Code)
	}
}
```
(imports as above; package `del`.) Note: this requires the dir's `NewServer` to also wire `GET /tickets/{id}` (working setup alongside the broken delete handler) — state that in the dir README.
READMEs: PathValue parsing + Atoi errors → 400; PUT semantics (full replace in this API); 204 No Content.

- [ ] **Step 3: `05_graceful` + `06_config`**

`shutdown_test.go`:
```go
package shutdown

import (
	"net/http"
	"testing"
	"time"
)

func TestShutdownGracefully(t *testing.T) {
	srv := &http.Server{}
	if err := ShutdownGracefully(srv, 2*time.Second); err != nil {
		t.Fatalf("ShutdownGracefully = %v, want nil", err)
	}
}
```

`shutdown.go` (imports all used in stub: errors via errors.New, http via signature, time via param):
```go
// Package shutdown teaches graceful shutdown.
package shutdown

import (
	"errors"
	"net/http"
	"time"
)

// ShutdownGracefully stops srv, waiting up to timeout.
// TODO: context.WithTimeout + srv.Shutdown (import context in the fix).
func ShutdownGracefully(srv *http.Server, timeout time.Duration) error {
	return errors.New("TODO")
}
```
README: Shutdown vs Close, draining with timeout (describe).

`config_test.go`:
```go
package config

import "testing"

func TestConfigFromEnv(t *testing.T) {
	t.Setenv("PORT", "9090")
	t.Setenv("DATA_FILE", "/tmp/t.json")
	got := ConfigFromEnv()
	if got.Port != 9090 || got.DataFile != "/tmp/t.json" {
		t.Fatalf("Config = %+v", got)
	}
}
```
Plus a defaults subtest that unsets? `t.Setenv` cannot unset; defaults covered by documenting + stub returns zeros (both fail pre-fix anyway). Keep single test (env set) — stub `Config{}` fails ✓. Also add unset-defaults test via `os.Unsetenv`? `t.Setenv` + `os.Unsetenv("PORT")` interplay is discouraged. Keep: test sets both vars; stub zeros fail. README documents defaults (8080, tickets.json).
```go
// Config holds server settings.
type Config struct {
	Port     int
	DataFile string
}

// ConfigFromEnv reads PORT and DATA_FILE, defaulting to 8080/tickets.json.
// TODO: os.Getenv + strconv.Atoi with fallback (import os, strconv).
func ConfigFromEnv() Config {
	return Config{}
}
```

- [ ] **Step 4: `07_e2e`**

`e2e_test.go`:
```go
package e2e

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func do(t *testing.T, method, url, body string) (int, string) {
	t.Helper()
	var rdr io.Reader
	if body != "" {
		rdr = strings.NewReader(body)
	}
	req, err := http.NewRequest(method, url, rdr)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, string(raw)
}

func TestEndToEnd(t *testing.T) {
	srv := httptest.NewServer(NewServer().Handler())
	defer srv.Close()
	code, body := do(t, http.MethodPost, srv.URL+"/tickets", `{"title":"A","status":"open"}`)
	if code != http.StatusCreated {
		t.Fatalf("POST = (%d, %s), want 201", code, body)
	}
	var created struct {
		ID int `json:"id"`
	}
	if err := json.Unmarshal([]byte(body), &created); err != nil || created.ID == 0 {
		t.Fatalf("POST body = %s, want id", body)
	}
	idPath := fmt.Sprintf("%s/tickets/%d", srv.URL, created.ID)
	if code, body := do(t, http.MethodGet, idPath, ""); code != http.StatusOK || !strings.Contains(body, `"title":"A"`) {
		t.Fatalf("GET = (%d, %s)", code, body)
	}
	if code, _ := do(t, http.MethodPut, idPath, `{"title":"B","status":"closed"}`); code != http.StatusOK {
		t.Fatalf("PUT = %d, want 200", code)
	}
	if code, _ := do(t, http.MethodDelete, idPath, ""); code != http.StatusNoContent {
		t.Fatalf("DELETE = %d, want 204", code)
	}
	if code, _ := do(t, http.MethodGet, idPath, ""); code != http.StatusNotFound {
		t.Fatalf("GET after delete = %d, want 404", code)
	}
}
```

`e2e.go`: FULL working store + handlers (setup — same shapes as `01_store_server`
through `04_delete`), but `NewServer` returns a server with an EMPTY mux:
```go
// NewServer builds the server. Handlers and store already work.
// TODO: register POST+GET /tickets and GET+PUT+DELETE /tickets/{id} on mux.
func NewServer() *Server {
	mux := http.NewServeMux()
	return &Server{store: &Store{}, mux: mux}
}

// Handler exposes the routes.
func (s *Server) Handler() http.Handler { return s.mux }
```
README: e2e anatomy, ephemeral ports, one broken layer at a time (describe).

- [ ] **Step 5: Verify and commit**

Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test ./exercises/15_capstone/...`
Expected: FAIL (7 packages; 501s/404s/empty wirings; e2e binds one ephemeral socket).

```bash
git add exercises/15_capstone
git commit -m "feat: add 15_capstone section (wave 3, 100 total)"
```

---

### Task 4: Extend `solutions` branch and release-check Wave 3 (100/100)

**Files:** none new (branch operation + stub fixes only; never edit tests or READMEs).

**Interfaces:**
- Consumes: all Wave 3 stubs + tests from Tasks 1–3.
- Produces: `solutions` branch where `go test ./...` passes 100/100.

- [ ] **Step 1: Bring `solutions` forward and apply fixes**

```bash
git checkout solutions
git merge main -m "merge: bring wave 3 stubs onto solutions"
```
(Trivial overlap class possible as in Wave 2; keep the `solutions` fixed version. Beyond trivial → STOP, NEEDS_CONTEXT.)
Apply minimal inverse-of-TODO fixes, one commit per new section:
```bash
git add exercises/13_http && git commit -m "solution: 13_http wave 3"
git add exercises/14_persistence && git commit -m "solution: 14_persistence wave 3"
git add exercises/15_capstone && git commit -m "solution: 15_capstone wave 3"
```

- [ ] **Step 2: Verify `solutions` is green**

Run (on `solutions`): `go vet ./... && go build ./... && go test -count=1 ./...`
Expected: PASS, 100/100 packages ok (15 sections: 5+15+8+7+5+6+4+7+6+4+5+6+10+5+7).

- [ ] **Step 3: Verify `main` and finish on `main`**

```bash
git checkout main
```
Run: `go vet ./... && go build ./...`
Expected: PASS.
Run: `go test ./...`
Expected: FAIL (22 new packages fail by design; 01_mutex still race-only per standing ruling).
Do NOT push.
