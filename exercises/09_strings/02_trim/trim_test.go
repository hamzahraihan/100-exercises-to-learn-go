package trim

import "testing"

func TestClean(t *testing.T) {
	if got, want := Clean("  hi\n"), "hi"; got != want {
		t.Fatalf("Clean = %q, want %q", got, want)
	}
}

func TestTrimExt(t *testing.T) {
	if got, want := TrimExt("notes.txt"), "notes"; got != want {
		t.Fatalf("TrimExt = %q, want %q", got, want)
	}
}
