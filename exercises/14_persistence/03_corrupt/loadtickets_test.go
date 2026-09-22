package loadj

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadTickets(t *testing.T) {
	dir := t.TempDir()
	ok := filepath.Join(dir, "ok.json")
	os.WriteFile(ok, []byte(`[{"id":1,"title":"a"}]`), 0o644)
	got, err := LoadTickets(ok)
	if err != nil || len(got) != 1 || got[0].Title != "a" {
		t.Fatalf("LoadTickets = (%+v, %v)", got, err)
	}
	bad := filepath.Join(dir, "bad.json")
	os.WriteFile(bad, []byte(`{oops`), 0o644)
	if _, err := LoadTickets(bad); err == nil {
		t.Fatal("corrupt file accepted, want descriptive error")
	} else if !strings.Contains(err.Error(), bad) {
		t.Fatalf("error = %q, want path inside", err.Error())
	}
}
