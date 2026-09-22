package exporttk

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestExport(t *testing.T) {
	path := filepath.Join(t.TempDir(), "backup.json")
	in := []Ticket{{ID: 1, Title: "a"}, {ID: 2, Title: "b"}}
	if err := Export(path, in); err != nil {
		t.Fatalf("Export errored: %v", err)
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("backup missing: %v", err)
	}
	var got []Ticket
	if err := json.Unmarshal(raw, &got); err != nil || len(got) != 2 || got[1].Title != "b" {
		t.Fatalf("backup = (%s, %v)", raw, err)
	}
}
