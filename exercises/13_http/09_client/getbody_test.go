package getbody

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("data"))
	}))
	defer srv.Close()
	got, err := GetBody(srv.URL)
	if err != nil || got != "data" {
		t.Fatalf("GetBody = (%q, %v), want (data, nil)", got, err)
	}
}
