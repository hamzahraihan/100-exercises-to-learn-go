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
