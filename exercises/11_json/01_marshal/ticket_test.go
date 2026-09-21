package ticket

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMarshalTicket(t *testing.T) {
	in := Ticket{ID: 1, Title: "Fix bug", Description: "Crash on login", Status: "open"}
	out, err := MarshalTicket(in)
	if err != nil {
		t.Fatalf("MarshalTicket errored: %v", err)
	}
	var got Ticket
	if err := json.Unmarshal([]byte(out), &got); err != nil {
		t.Fatalf("output is not valid JSON: %v (output = %q)", err, out)
	}
	if !reflect.DeepEqual(got, in) {
		t.Fatalf("round-trip = %+v, want %+v", got, in)
	}
}
