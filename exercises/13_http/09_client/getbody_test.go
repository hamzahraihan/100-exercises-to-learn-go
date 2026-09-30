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

	if _, err := GetBody("http://127.0.0.1:0"); err == nil {
		t.Fatal("GetBody(bad URL) = nil error, want connection error")
	}

	empty := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer empty.Close()
	got, err = GetBody(empty.URL)
	if err != nil || got != "" {
		t.Fatalf("GetBody(empty) = (%q, %v), want (\"\", nil)", got, err)
	}
}
