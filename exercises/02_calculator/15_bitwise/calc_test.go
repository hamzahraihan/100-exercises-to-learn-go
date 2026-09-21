package calc

import "testing"

func TestHas(t *testing.T) {
	if !Has(Read|Write, Write) {
		t.Fatal("Has(Read|Write, Write) = false, want true")
	}
	if Has(Read, Write) {
		t.Fatal("Has(Read, Write) = true, want false")
	}
	if !Has(Execute, Execute) {
		t.Fatal("Has(Execute, Execute) = false, want true")
	}
}
