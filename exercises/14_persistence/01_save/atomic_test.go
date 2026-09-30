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

	// Overwrite must replace content atomically and leave no temp behind.
	if err := Save(path, []byte(`{"id":2}`)); err != nil {
		t.Fatalf("Save overwrite errored: %v", err)
	}
	got, err = os.ReadFile(path)
	if err != nil || string(got) != `{"id":2}` {
		t.Fatalf("overwrite file = (%q, %v)", got, err)
	}
	ents, err := os.ReadDir(filepath.Dir(path))
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range ents {
		if e.Name() != "data.json" {
			t.Fatalf("stray file after Save: %q", e.Name())
		}
	}
}
