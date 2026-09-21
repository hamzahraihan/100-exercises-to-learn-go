package store

import (
	"reflect"
	"testing"
)

func TestCloneTickets(t *testing.T) {
	in := []Ticket{{ID: 1, Title: "a"}, {ID: 2, Title: "b"}}
	out := CloneTickets(in)
	if !reflect.DeepEqual(out, in) {
		t.Fatalf("CloneTickets = %v, want %v", out, in)
	}
	out[0].Title = "changed"
	if in[0].Title != "a" {
		t.Fatalf("clone aliases original: %v", in)
	}
}
