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
