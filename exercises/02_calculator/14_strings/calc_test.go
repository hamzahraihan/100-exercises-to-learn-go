package calc

import "testing"

func TestToUpperFirst(t *testing.T) {
	if got, want := ToUpperFirst("hello"), "Hello"; got != want {
		t.Fatalf("ToUpperFirst(%q) = %q, want %q", "hello", got, want)
	}
	if got := ToUpperFirst(""); got != "" {
		t.Fatalf("ToUpperFirst(%q) = %q, want %q", "", got, "")
	}
}
