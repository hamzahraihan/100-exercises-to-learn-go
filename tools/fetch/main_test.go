package main

import (
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
