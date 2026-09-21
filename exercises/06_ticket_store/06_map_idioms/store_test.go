package store

import "testing"

func TestRegistry(t *testing.T) {
	var r Registry
	r.Put(Ticket{Title: "a"})
	tk, ok := r.Lookup("a")
	if !ok || tk.Title != "a" {
		t.Fatalf("Lookup(%q) = (%+v, %v), want found", "a", tk, ok)
	}
	if _, ok := r.Lookup("missing"); ok {
		t.Fatal("Lookup(missing) found, want false")
	}
}
