package atomic

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSave(t *testing.T) {
	path := filepath.Join(t.TempDir(), "data.json")
	if err := Save(path, []byte(`{"id":1}`)); err != nil {
		t.Fatalf("Save errored: %v", err)
	}
	got, err := os.ReadFile(path)
	if err != nil || string(got) != `{"id":1}` {
		t.Fatalf("file = (%q, %v)", got, err)
	}
}
