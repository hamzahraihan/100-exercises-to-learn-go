package words

import (
	"reflect"
	"testing"
)

func TestWords(t *testing.T) {
	want := []string{"go", "is", "fun"}
	if got := Words("  go  is fun "); !reflect.DeepEqual(got, want) {
		t.Fatalf("Words = %v, want %v", got, want)
	}
}

func TestJoinWords(t *testing.T) {
	if got, want := JoinWords([]string{"a", "b"}), "a b"; got != want {
		t.Fatalf("JoinWords = %q, want %q", got, want)
	}
}
