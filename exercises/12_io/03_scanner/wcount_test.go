package wcount

import (
	"strings"
	"testing"
)

func TestWordCount(t *testing.T) {
	if got, _ := WordCount(strings.NewReader("go is fun")); got != 3 {
		t.Fatalf("WordCount = %d, want 3", got)
	}
	if got, err := WordCount(strings.NewReader("")); got != 0 || err != nil {
		t.Fatalf("WordCount(empty) = (%d, %v), want (0, nil)", got, err)
	}
}
