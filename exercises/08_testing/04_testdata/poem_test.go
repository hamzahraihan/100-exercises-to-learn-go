package poem

import "testing"

func TestFirstLine(t *testing.T) {
	got, err := FirstLine("testdata/poem.txt")
	if err != nil {
		t.Fatalf("FirstLine errored: %v", err)
	}
	if got != "Roses are red" {
		t.Fatalf("FirstLine = %q, want %q", got, "Roses are red")
	}
}
