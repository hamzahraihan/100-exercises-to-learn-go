package ticket

import "testing"

func TestUnmarshalTicket(t *testing.T) {
	got, err := UnmarshalTicket(`{"id":1,"title":"Fix bug","status":"open"}`)
	if err != nil {
		t.Fatalf("valid JSON rejected: %v", err)
	}
	if got.Title != "Fix bug" || got.ID != 1 {
		t.Fatalf("decoded = %+v", got)
	}
	if _, err := UnmarshalTicket(`{"id":1,"status":"open"}`); err == nil {
		t.Fatal("missing title accepted, want validation error")
	}
}
