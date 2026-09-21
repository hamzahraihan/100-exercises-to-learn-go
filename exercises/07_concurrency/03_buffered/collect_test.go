package collect

import (
	"reflect"
	"testing"
)

func TestCollect(t *testing.T) {
	if got, want := Collect(5), []int{0, 2, 4, 6, 8}; !reflect.DeepEqual(got, want) {
		t.Fatalf("Collect(5) = %v, want %v", got, want)
	}
	if got, want := Collect(1), []int{0}; !reflect.DeepEqual(got, want) {
		t.Fatalf("Collect(1) = %v, want %v", got, want)
	}
}
