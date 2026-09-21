package status

import "testing"

func TestStatusValid(t *testing.T) {
	for _, s := range []Status{StatusOpen, StatusInProgress, StatusClosed} {
		if !s.Valid() {
			t.Fatalf("Status(%d).Valid() = false, want true", int(s))
		}
	}
	if Status(99).Valid() {
		t.Fatal("Status(99).Valid() = true, want false")
	}
}
