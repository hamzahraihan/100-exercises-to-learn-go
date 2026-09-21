package pipeline

import (
	"reflect"
	"testing"
)

func TestDoubleAll(t *testing.T) {
	got := DoubleAll([]int{1, 2, 3, 4})
	if !reflect.DeepEqual(got, []int{2, 4, 6, 8}) {
		t.Fatalf("DoubleAll = %v, want [2 4 6 8]", got)
	}
}
