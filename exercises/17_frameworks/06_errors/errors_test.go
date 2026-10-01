package gerrors

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() { gin.SetMode(gin.TestMode) }

func TestErrorEnvelope(t *testing.T) {
	r := NewRouter()
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/tickets/7", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("GET /tickets/7 = %d, want 200", rec.Code)
	}
	rec2 := httptest.NewRecorder()
	r.ServeHTTP(rec2, httptest.NewRequest(http.MethodGet, "/tickets/99", nil))
	if rec2.Code != http.StatusNotFound {
		t.Fatalf("GET /tickets/99 = %d, want 404", rec2.Code)
	}
	if !strings.Contains(rec2.Body.String(), `"error"`) {
		t.Fatalf("body = %s, want error envelope", rec2.Body.String())
	}
}
