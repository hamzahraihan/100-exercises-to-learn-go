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
