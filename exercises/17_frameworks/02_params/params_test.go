package gparams

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() { gin.SetMode(gin.TestMode) }

func TestParams(t *testing.T) {
	r := NewRouter()
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/tickets/7", nil))
	if rec.Body.String() != "one:7" {
		t.Fatalf("GET /tickets/7 = %q, want one:7", rec.Body.String())
	}
	rec2 := httptest.NewRecorder()
	r.ServeHTTP(rec2, httptest.NewRequest(http.MethodGet, "/tickets", nil))
	if rec2.Code != http.StatusNotFound {
		t.Fatalf("GET /tickets = %d, want 404", rec2.Code)
	}
}
