package ticket

import "testing"

func TestNewTicket(t *testing.T) {
	tk := NewTicket("Fix bug", "Crash on login")
	if tk.Title != "Fix bug" || tk.Description != "Crash on login" {
		t.Fatalf("got %+v", tk)
	}
}
