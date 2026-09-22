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
