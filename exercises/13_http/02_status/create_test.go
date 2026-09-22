package created

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	Create(rec, httptest.NewRequest(http.MethodPost, "/tickets", nil))
	if rec.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusCreated)
	}
}
