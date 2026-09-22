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
