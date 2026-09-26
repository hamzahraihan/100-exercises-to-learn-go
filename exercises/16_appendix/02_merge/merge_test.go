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
}
