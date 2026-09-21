package keys

import (
	"reflect"
	"testing"
)

func TestKeys(t *testing.T) {
	if got, want := Keys(map[string]int{"a": 1}), []string{"a"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("Keys = %v, want %v", got, want)
	}
	if got := Keys(map[int]string{}); len(got) != 0 {
		t.Fatalf("Keys of empty map = %v, want empty", got)
	}
}
