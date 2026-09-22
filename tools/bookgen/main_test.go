package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseWeight(t *testing.T) {
	got, err := parseWeight("01_syntax")
	if err != nil || got != 1 {
		t.Fatalf("parseWeight(01_syntax) = %d, %v; want 1, nil", got, err)
	}
	if _, err := parseWeight("nosuchprefix"); err == nil {
		t.Fatal("expected error for missing numeric prefix")
	}
}

func TestExtractTitle(t *testing.T) {
	got, err := extractTitle("# Syntax\n\nBody here.\n")
	if err != nil || got != "Syntax" {
		t.Fatalf("extractTitle = %q, %v; want %q", got, err, "Syntax")
	}
	if _, err := extractTitle("no heading here\n"); err == nil {
		t.Fatal("expected error for missing H1")
	}
}

func TestEmitQuotesTitleAndFooter(t *testing.T) {
	ex := exercise{Section: "01_intro", Name: "01_syntax", Title: `A "quoted": title`, Body: "# T\n\nBody.\r\n"}
	page := emit(ex)
	if !strings.Contains(page, `title: "A \"quoted\": title"`) {
		t.Fatalf("front-matter quoting missing:\n%s", page)
	}
	if strings.Contains(page, "\r") {
		t.Fatal("CRLF leaked into output")
	}
	if !strings.Contains(page, "exercises/01_intro/01_syntax/README.md") {
		t.Fatal("source-backlink footer missing")
	}
	if !strings.Contains(page, "Mainmatter") || !strings.Contains(page, "CC BY-NC 4.0") {
		t.Fatal("attribution footer missing")
	}
	if !strings.HasPrefix(page, "---\n# DO NOT EDIT") {
		t.Fatal("generated header missing inside front-matter")
	}
}

func TestWalkExercisesDuplicateWeight(t *testing.T) {
	root := t.TempDir()
	for _, dir := range []string{"01_a", "01_b"} {
		p := filepath.Join(root, "01_sec", dir)
		if err := os.MkdirAll(p, 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(p, "README.md"), []byte("# T\n"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := walkExercises(root); err == nil {
		t.Fatal("expected duplicate-weight error, got nil")
	}
}
