package importtk

import (
	"os"
	"path/filepath"
	"testing"
)

func TestImport(t *testing.T) {
	dir := t.TempDir()
	ok := filepath.Join(dir, "ok.json")
	os.WriteFile(ok, []byte(`[{"id":1,"title":"a"}]`), 0o644)
	got, err := Import(ok)
	if err != nil || len(got) != 1 || got[0].Title != "a" {
		t.Fatalf("Import = (%+v, %v)", got, err)
	}
	bad := filepath.Join(dir, "bad.json")
	os.WriteFile(bad, []byte(`[{"id":1,"title":""}]`), 0o644)
	if _, err := Import(bad); err == nil {
		t.Fatal("empty title imported, want validation error")
	}
}
