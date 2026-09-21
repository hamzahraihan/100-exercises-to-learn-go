package parse

import "testing"

func TestSumStrings(t *testing.T) {
	if got, err := SumStrings("3", "4"); err != nil || got != 7 {
		t.Fatalf("SumStrings = (%d, %v), want (7, nil)", got, err)
	}
	if _, err := SumStrings("x", "4"); err == nil {
		t.Fatal("SumStrings(x) accepted, want error")
	}
}
