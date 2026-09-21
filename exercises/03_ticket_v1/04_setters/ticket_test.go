package ticket

import "testing"

func TestSetTitle(t *testing.T) {
	tk := Ticket{Title: "old", Description: "long enough description"}
	tk.SetTitle("new")
	if tk.Title != "new" {
		t.Fatalf("Title = %q, want %q", tk.Title, "new")
	}
}
