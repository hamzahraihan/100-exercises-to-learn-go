package must

import "testing"

func TestMustParse(t *testing.T) {
	if got := MustParse("hi"); got != 2 {
		t.Fatalf("MustParse(%q) = %d, want 2", "hi", got)
	}
	if got := MustParse(""); got != -1 {
		t.Fatalf("MustParse(%q) = %d, want -1 (no panic must escape)", "", got)
	}
}
