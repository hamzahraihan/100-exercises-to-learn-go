package greet

import "testing"

func TestShout(t *testing.T) {
	if got, want := Shout("hello"), "HELLO"; got != want {
		t.Fatalf("Shout(%q) = %q, want %q", "hello", got, want)
	}
}
