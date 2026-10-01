package gbind

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() { gin.SetMode(gin.TestMode) }

func TestBindHappy(t *testing.T) {
	r := NewRouter()
	body := strings.NewReader(`{"id":1,"title":"Fix bug","status":"open"}`)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/tickets", body))
	if rec.Code != http.StatusCreated {
		t.Fatalf("status = %d, want 201", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), `"title":"Fix bug"`) {
		t.Fatalf("body = %s, want title echoed", rec.Body.String())
	}
}

func TestBindBad(t *testing.T) {
	r := NewRouter()
	body := strings.NewReader(`{"id":1,"status":"open"}`)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/tickets", body))
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", rec.Code)
	}
}
