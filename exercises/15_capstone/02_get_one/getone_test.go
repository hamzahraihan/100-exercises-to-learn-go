package getone

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

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
