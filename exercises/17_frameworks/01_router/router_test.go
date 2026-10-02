package grouter

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() { gin.SetMode(gin.TestMode) }

func get(t *testing.T, r http.Handler, path string) *httptest.ResponseRecorder {
	t.Helper()
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
	return rec
}

func TestRouter(t *testing.T) {
	r := NewRouter()
	if rec := get(t, r, "/tickets"); rec.Body.String() != "list" {
		t.Fatalf("GET /tickets = %q, want list", rec.Body.String())
	}
	if rec := get(t, r, "/nope"); rec.Code != http.StatusNotFound {
		t.Fatalf("GET /nope = %d, want 404", rec.Code)
	}
}
