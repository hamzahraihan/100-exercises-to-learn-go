package store

import (
	"reflect"
	"testing"
)

func TestSortedByID(t *testing.T) {
	in := []Ticket{{ID: 3}, {ID: 1}, {ID: 2}}
	want := []Ticket{{ID: 1}, {ID: 2}, {ID: 3}}
	if got := SortedByID(in); !reflect.DeepEqual(got, want) {
		t.Fatalf("SortedByID = %v, want %v", got, want)
	}
}

func TestOpenOnly(t *testing.T) {
	in := []Ticket{{ID: 1, Closed: true}, {ID: 2}, {ID: 3, Closed: true}}
	want := []Ticket{{ID: 2}}
	if got := OpenOnly(in); !reflect.DeepEqual(got, want) {
		t.Fatalf("OpenOnly = %v, want %v", got, want)
	}
}
