package store

import "testing"

func TestMapStore(t *testing.T) {
	var s Store
	id := s.Add("first")
	tk, ok := s.Get(id)
	if !ok || tk.Title != "first" {
		t.Fatalf("Get(%d) = %+v, %v", id, tk, ok)
	}
	if _, ok := s.Get(999); ok {
		t.Fatal("Get(999) found ticket, want missing")
	}
}
