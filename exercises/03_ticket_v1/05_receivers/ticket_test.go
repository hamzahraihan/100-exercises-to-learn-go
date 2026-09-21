package ticket

import "testing"

func TestRenamed(t *testing.T) {
	orig := Ticket{Title: "a", Description: "long enough description"}
	got := orig.Renamed("b")
	if got.Title != "b" || got.Description != orig.Description {
		t.Fatalf("Renamed = %+v, want title b with description kept", got)
	}
	if orig.Title != "a" {
		t.Fatalf("original mutated: %+v", orig)
	}
}
