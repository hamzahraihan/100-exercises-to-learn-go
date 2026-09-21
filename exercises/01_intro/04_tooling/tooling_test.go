package tooling

import "testing"

func TestQuote(t *testing.T) {
	if got, want := Quote("hi"), `"hi"`; got != want {
		t.Fatalf("Quote(%q) = %q, want %q", "hi", got, want)
	}
}
