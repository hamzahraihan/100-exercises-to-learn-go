package ticket

import "testing"

func TestIsZero(t *testing.T) {
	if !(Ticket{}).IsZero() {
		t.Fatal("Ticket{}.IsZero() = false, want true")
	}
	if (Ticket{Title: "x"}).IsZero() {
		t.Fatal("non-empty ticket reported zero")
	}
}
