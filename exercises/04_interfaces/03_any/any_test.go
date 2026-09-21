package show

import "testing"

func TestSprintAny(t *testing.T) {
	if got, want := SprintAny(42), "42"; got != want {
		t.Fatalf("SprintAny(42) = %q, want %q", got, want)
	}
	if got, want := SprintAny("hi"), "hi"; got != want {
		t.Fatalf("SprintAny(%q) = %q, want %q", "hi", got, want)
	}
}
