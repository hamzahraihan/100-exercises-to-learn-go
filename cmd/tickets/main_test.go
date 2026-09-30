package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestConfigFromEnv(t *testing.T) {
	t.Setenv("PORT", "")
	t.Setenv("DATA_FILE", "")
	if got := ConfigFromEnv(); got.Port != 8080 || got.DataFile != "tickets.json" {
		t.Fatalf("defaults = %+v", got)
	}
	for _, tc := range []struct {
		in   string
		want int
	}{
		{"8080", 8080}, {"3000", 3000}, {"1", 1}, {"65535", 65535},
		{"", 8080}, {"abc", 8080}, {"0", 8080}, {"-1", 8080}, {"99999", 8080},
	} {
		t.Setenv("PORT", tc.in)
		if got := ConfigFromEnv(); got.Port != tc.want {
			t.Errorf("PORT=%q => %d, want %d", tc.in, got.Port, tc.want)
		}
	}
	t.Setenv("PORT", "8080")
	t.Setenv("DATA_FILE", "custom.json")
	if got := ConfigFromEnv(); got.DataFile != "custom.json" {
		t.Fatalf("DATA_FILE = %q", got.DataFile)
	}
}

func TestStoreCRUD(t *testing.T) {
	s := NewStore([]Ticket{{ID: 5, Title: "old"}}, "")
	got, err := s.Add(Ticket{Title: "new"})
	if err != nil {
		t.Fatal(err)
	}
	if got.ID != 6 {
		t.Fatalf("Add ID = %d, want 6", got.ID)
	}
	if _, ok := s.Get(5); !ok {
		t.Fatal("Get(5) missing")
	}
	updated, ok, err := s.Update(5, Ticket{Title: "upd"})
	if err != nil {
		t.Fatal(err)
	}
	if !ok || updated.ID != 5 || updated.Title != "upd" {
		t.Fatalf("Update = %+v, %v", updated, ok)
	}
	if _, ok, _ := s.Update(999, Ticket{Title: "x"}); ok {
		t.Fatal("Update missing should fail")
	}
	list := s.List()
	if len(list) != 2 {
		t.Fatalf("List len = %d", len(list))
	}
	list[0].Title = "mutated"
	if fresh, _ := s.Get(list[0].ID); fresh.Title == "mutated" {
		t.Fatal("List must return a copy")
	}
	if ok, _ := s.Delete(5); !ok {
		t.Fatal("Delete should succeed")
	}
	if ok, _ := s.Delete(5); ok {
		t.Fatal("second Delete should fail")
	}
}

func TestSaveLoadRoundTrip(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "tickets.json")
	ts := []Ticket{{ID: 1, Title: "a", Status: "open"}, {ID: 2, Title: "b", Status: "closed"}}
	if err := saveTickets(p, ts); err != nil {
		t.Fatal(err)
	}
	back, err := loadTickets(p)
	if err != nil || len(back) != 2 || back[0].Title != "a" {
		t.Fatalf("round-trip = %+v, %v", back, err)
	}
	if got, err := loadTickets(filepath.Join(dir, "missing.json")); err != nil || got != nil {
		t.Fatalf("missing = %+v, %v", got, err)
	}
	os.WriteFile(filepath.Join(dir, "bad.json"), []byte("{invalid"), 0o644)
	if _, err := loadTickets(filepath.Join(dir, "bad.json")); err == nil {
		t.Fatal("corrupt should error")
	}
}

func doReq(s *Server, method, target, body string) *httptest.ResponseRecorder {
	var r *strings.Reader
	if body == "" {
		r = strings.NewReader("")
	} else {
		r = strings.NewReader(body)
	}
	req := httptest.NewRequest(method, target, r)
	w := httptest.NewRecorder()
	s.mux.ServeHTTP(w, req)
	return w
}

func TestHandlers(t *testing.T) {
	s := NewServer(NewStore(nil, filepath.Join(t.TempDir(), "tickets.json")))

	w := doReq(s, "POST", "/tickets", `{"title":""}`)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("empty title = %d", w.Code)
	}
	w = doReq(s, "POST", "/tickets", `{"title":"Fix bug","status":"open"}`)
	if w.Code != http.StatusCreated {
		t.Fatalf("create = %d %s", w.Code, w.Body.String())
	}
	var created Ticket
	if err := json.Unmarshal(w.Body.Bytes(), &created); err != nil || created.ID == 0 {
		t.Fatalf("created = %s, %v", w.Body.String(), err)
	}
	if w := doReq(s, "GET", "/tickets", ""); w.Code != http.StatusOK {
		t.Fatalf("list = %d", w.Code)
	}
	if w := doReq(s, "GET", "/tickets/abc", ""); w.Code != http.StatusBadRequest {
		t.Fatalf("bad id = %d", w.Code)
	}
	if w := doReq(s, "GET", "/tickets/9999", ""); w.Code != http.StatusNotFound {
		t.Fatalf("missing = %d", w.Code)
	}
	w = doReq(s, "PUT", "/tickets/abc", `{"title":"x"}`)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("put bad id = %d", w.Code)
	}
	big := `{"title":"` + strings.Repeat("x", (1<<20)+10) + `"}`
	if w := doReq(s, "POST", "/tickets", big); w.Code != http.StatusBadRequest {
		t.Fatalf("oversize = %d, want 400", w.Code)
	}
}

func TestConcurrentMutationsPersist(t *testing.T) {
	p := filepath.Join(t.TempDir(), "tickets.json")
	s := NewServer(NewStore(nil, p))
	const n = 20
	done := make(chan int, n)
	for i := 0; i < n; i++ {
		go func() {
			w := doReq(s, "POST", "/tickets", `{"title":"t"}`)
			done <- w.Code
		}()
	}
	for i := 0; i < n; i++ {
		if code := <-done; code != http.StatusCreated {
			t.Fatalf("concurrent create = %d", code)
		}
	}
	back, err := loadTickets(p)
	if err != nil {
		t.Fatal(err)
	}
	if len(back) != n {
		t.Fatalf("file has %d tickets, want %d", len(back), n)
	}
	seen := map[int]bool{}
	for _, tk := range back {
		if seen[tk.ID] {
			t.Fatalf("duplicate ID %d", tk.ID)
		}
		seen[tk.ID] = true
	}
}

func TestNewHTTPServerAndShutdown(t *testing.T) {
	srv := NewServer(NewStore(nil, ""))
	httpSrv := newHTTPServer(Config{Port: 8080, DataFile: "tickets.json"}, srv)
	if httpSrv.Addr != ":8080" {
		t.Fatalf("Addr = %q", httpSrv.Addr)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := waitAndShutdown(ctx, httpSrv, time.Second); err != nil {
		t.Fatalf("shutdown = %v", err)
	}
}
