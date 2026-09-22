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
