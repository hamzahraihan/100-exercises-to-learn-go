package calc

import "testing"

func TestAddUint8(t *testing.T) {
	if got, overflow := AddUint8(100, 100); got != 200 || overflow {
		t.Fatalf("AddUint8(100,100) = (%d,%v), want (200,false)", got, overflow)
	}
	if _, overflow := AddUint8(200, 100); !overflow {
		t.Fatal("AddUint8(200,100) should report overflow")
	}
}
