package store

import "testing"

func fiveTickets() []Ticket {
	return []Ticket{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}}
}

func TestPage(t *testing.T) {
	ts := fiveTickets()
	got := Page(ts, 1, 2)
	if len(got) != 2 || got[0].ID != 2 || got[1].ID != 3 {
		t.Fatalf("Page(offset 1, limit 2) = %v, want IDs [2 3]", got)
	}
	if got := Page(ts, 99, 10); len(got) != 0 {
		t.Fatalf("Page past end = %v, want empty", got)
	}
	if got := Page(ts, 0, 0); len(got) != 0 {
		t.Fatalf("Page limit 0 = %v, want empty", got)
	}
	if got := Page(ts, -1, -5); len(got) != 0 {
		t.Fatalf("Page negative = %v, want empty (clamped, no panic)", got)
	}
}
