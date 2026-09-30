package merge

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	a := []Ticket{{ID: 1, Title: "A"}, {ID: 2, Title: "B"}}
	b := []Ticket{{ID: 2, Title: "B2"}, {ID: 3, Title: "C"}}
	want := []Ticket{{ID: 1, Title: "A"}, {ID: 2, Title: "B2"}, {ID: 3, Title: "C"}}
	if got := Merge(a, b); !reflect.DeepEqual(got, want) {
		t.Fatalf("Merge = %v, want %v", got, want)
	}
	if got := Merge([]Ticket{{ID: 1, Title: "A"}}, []Ticket{{ID: 2, Title: "B"}}); len(got) != 2 {
		t.Fatalf("Merge disjoint = %v, want both tickets", got)
	}

	if got := Merge(nil, nil); len(got) != 0 {
		t.Fatalf("Merge(nil,nil) = %v, want empty", got)
	}
	if got := Merge(nil, []Ticket{{ID: 9, Title: "N"}}); len(got) != 1 || got[0].ID != 9 {
		t.Fatalf("Merge(nil,b) = %v, want b's ticket", got)
	}

	// Inputs must not be mutated; duplicate IDs in b are last-wins, no dupes.
	aa := []Ticket{{ID: 1, Title: "A"}}
	bb := []Ticket{{ID: 1, Title: "B1"}, {ID: 1, Title: "B2"}}
	got := Merge(aa, bb)
	if len(got) != 1 || got[0].Title != "B2" {
		t.Fatalf("Merge dup = %v, want single B2", got)
	}
	if aa[0].Title != "A" || bb[0].Title != "B1" {
		t.Fatalf("inputs mutated: a=%v b=%v", aa, bb)
	}
}
