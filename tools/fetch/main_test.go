package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScanFindsSyntax(t *testing.T) {
	list, err := scanExercises(filepath.Join("..", "..", "exercises"))
	if err != nil {
		t.Fatalf("scanExercises = %v", err)
	}
	found := false
	for _, e := range list {
		if e.Section == "01_intro" && e.Name == "01_syntax" {
			found = true
			if len(e.Files) != 3 {
				t.Fatalf("Files = %v, want 3 entries", e.Files)
			}
		}
	}
	if !found {
		t.Fatal("01_intro/01_syntax not in scan")
	}
}

func TestResolveExactAndAmbiguous(t *testing.T) {
	list := []entry{
		{Section: "01_intro", Name: "01_syntax"},
		{Section: "01_intro", Name: "02_test_workflow"},
	}
	got, err := resolve("01_syntax", list)
	if err != nil || got.Name != "01_syntax" {
		t.Fatalf("resolve exact = %v, %v", got, err)
	}
	got, err = resolve("01_intro/01_syntax", list)
	if err != nil || got.Section != "01_intro" {
		t.Fatalf("resolve qualified = %v, %v", got, err)
	}
	if _, err := resolve("01_", list); err == nil {
		t.Fatal("expected ambiguous error for '01_'")
	}
	if _, err := resolve("nope_xyz", list); err == nil {
		t.Fatal("expected not-found error")
	}
}

func TestCopyAndGoMod(t *testing.T) {
	src := t.TempDir()
	exDir := filepath.Join(src, "01_intro", "01_syntax")
	if err := os.MkdirAll(exDir, 0o755); err != nil {
		t.Fatal(err)
	}
	os.WriteFile(filepath.Join(exDir, "README.md"), []byte("# T\n"), 0o644)
	os.WriteFile(filepath.Join(exDir, "syntax.go"), []byte("package syntax\n"), 0o644)
	os.WriteFile(filepath.Join(exDir, "syntax_test.go"), []byte("package syntax\n"), 0o644)
	e := entry{Section: "01_intro", Name: "01_syntax", Files: []string{"README.md", "syntax.go", "syntax_test.go"}}
	out := filepath.Join(t.TempDir(), "out")
	if err := copyExercise(src, out, e); err != nil {
		t.Fatal(err)
	}
	if err := writeGoMod(out, "01_syntax"); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(out, "go.mod")); err != nil {
		t.Fatal("go.mod missing")
	}
}
