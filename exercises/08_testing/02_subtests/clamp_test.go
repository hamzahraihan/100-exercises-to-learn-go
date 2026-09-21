package clamp

import "testing"

func TestClamp(t *testing.T) {
	t.Run("above", func(t *testing.T) {
		if got, want := Clamp(99, 0, 10), 10; got != want {
			t.Fatalf("Clamp(99, 0, 10) = %d, want %d", got, want)
		}
	})
	t.Run("below", func(t *testing.T) {
		if got, want := Clamp(-5, 0, 10), 0; got != want {
			t.Fatalf("Clamp(-5, 0, 10) = %d, want %d", got, want)
		}
	})
	t.Run("inside", func(t *testing.T) {
		if got, want := Clamp(4, 0, 10), 4; got != want {
			t.Fatalf("Clamp(4, 0, 10) = %d, want %d", got, want)
		}
	})
}
