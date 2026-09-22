package del

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

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
