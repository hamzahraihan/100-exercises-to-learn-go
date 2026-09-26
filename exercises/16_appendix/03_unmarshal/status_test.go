package status

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalStatus(t *testing.T) {
	var tk Ticket
	if err := json.Unmarshal([]byte(`{"id":1,"title":"t","status":"closed"}`), &tk); err != nil {
		t.Fatalf("valid status rejected: %v", err)
	}
	if tk.Status != StatusClosed {
		t.Fatalf("Status = %d, want StatusClosed", tk.Status)
	}
	var bad Ticket
	if err := json.Unmarshal([]byte(`{"id":1,"title":"t","status":"bogus"}`), &bad); err == nil {
		t.Fatal("bogus status accepted, want error")
	}
}
