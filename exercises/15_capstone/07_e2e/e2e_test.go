package e2e

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func do(t *testing.T, method, url, body string) (int, string) {
	t.Helper()
	var rdr io.Reader
	if body != "" {
		rdr = strings.NewReader(body)
	}
	req, err := http.NewRequest(method, url, rdr)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, string(raw)
}

func TestEndToEnd(t *testing.T) {
	srv := httptest.NewServer(NewServer().Handler())
	defer srv.Close()
	code, body := do(t, http.MethodPost, srv.URL+"/tickets", `{"title":"A","status":"open"}`)
	if code != http.StatusCreated {
		t.Fatalf("POST = (%d, %s), want 201", code, body)
	}
	var created struct {
		ID int `json:"id"`
	}
	if err := json.Unmarshal([]byte(body), &created); err != nil || created.ID == 0 {
		t.Fatalf("POST body = %s, want id", body)
	}
	idPath := fmt.Sprintf("%s/tickets/%d", srv.URL, created.ID)
	if code, body := do(t, http.MethodGet, idPath, ""); code != http.StatusOK || !strings.Contains(body, `"title":"A"`) {
		t.Fatalf("GET = (%d, %s)", code, body)
	}
	if code, _ := do(t, http.MethodPut, idPath, `{"title":"B","status":"closed"}`); code != http.StatusOK {
		t.Fatalf("PUT = %d, want 200", code)
	}
	if code, _ := do(t, http.MethodDelete, idPath, ""); code != http.StatusNoContent {
		t.Fatalf("DELETE = %d, want 204", code)
	}
	if code, _ := do(t, http.MethodGet, idPath, ""); code != http.StatusNotFound {
		t.Fatalf("GET after delete = %d, want 404", code)
	}
}
