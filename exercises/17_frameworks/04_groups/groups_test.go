package ggroups

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() { gin.SetMode(gin.TestMode) }

func TestGroups(t *testing.T) {
	r := NewRouter()
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/api/v1/tickets", nil))
	if rec.Body.String() != "list" {
		t.Fatalf("GET /api/v1/tickets = %q, want list", rec.Body.String())
	}
	rec2 := httptest.NewRecorder()
	r.ServeHTTP(rec2, httptest.NewRequest(http.MethodGet, "/tickets", nil))
	if rec2.Code != http.StatusNotFound {
		t.Fatalf("GET /tickets = %d, want 404", rec2.Code)
	}
}
