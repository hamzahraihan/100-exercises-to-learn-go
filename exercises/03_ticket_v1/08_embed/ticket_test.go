package ticket

import "testing"

func TestNewTicketWithID(t *testing.T) {
	tk := NewTicketWithID(7, "bug")
	if tk.ID != 7 || tk.Title != "bug" {
		t.Fatalf("got %+v, want ID 7 title bug", tk)
	}
}
