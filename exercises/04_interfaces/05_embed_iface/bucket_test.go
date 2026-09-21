package bucket

import "testing"

func TestReadWriter(t *testing.T) {
	var b Bucket
	b.Write("hi")
	if got := b.Read(); got != "hi" {
		t.Fatalf("Read() = %q, want %q", got, "hi")
	}
}
