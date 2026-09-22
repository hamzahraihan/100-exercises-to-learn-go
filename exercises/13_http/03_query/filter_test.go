package filter

import (
	"net/http/httptest"
	"testing"
)

func TestStatusParam(t *testing.T) {
	r := httptest.NewRequest("GET", "/tickets?status=open", nil)
	if got, want := StatusParam(r), "open"; got != want {
		t.Fatalf("StatusParam = %q, want %q", got, want)
	}
	r = httptest.NewRequest("GET", "/tickets", nil)
	if got, want := StatusParam(r), "all"; got != want {
		t.Fatalf("StatusParam = %q, want %q", got, want)
	}
}
