package update

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

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
