package loadst

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "data.json")
	if err := os.WriteFile(path, []byte("hi"), 0o644); err != nil {
		t.Fatal(err)
	}
	if got, err := Load(path); err != nil || string(got) != "hi" {
		t.Fatalf("Load = (%q, %v)", got, err)
	}
	if got, err := Load(filepath.Join(dir, "missing.json")); err != nil || got != nil {
		t.Fatalf("Load(missing) = (%q, %v), want (nil, nil)", got, err)
	}
}
