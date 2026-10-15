package benchallocs

import (
	"testing"
)

func TestConcatFastCorrectness(t *testing.T) {
	cases := []struct {
		s    string
		n    int
		want string
	}{
		{"ab", 3, "ababab"},
		{"", 5, ""},
		{"x", 0, ""},
	}
	for _, c := range cases {
		got := ConcatFast(c.s, c.n)
		if got != c.want {
			t.Fatalf("ConcatFast(%q, %d) = %q, want %q (hint: build the result with strings.Builder)", c.s, c.n, got, c.want)
		}
	}
}

func TestAllocsPerRun(t *testing.T) {
	got := testing.AllocsPerRun(100, func() {
		ConcatFast("hello", 50)
	})
	if got > 1.0 {
		t.Fatalf("got %.1f allocs/op, want <= 1 (hint: use strings.Builder with b.Grow(len(s)*n) before the write loop)", got)
	}
}

func BenchmarkConcatSlow(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		ConcatSlow("ab", 100)
	}
}

func BenchmarkConcatFast(b *testing.B) {
	b.ReportAllocs()
	// TODO: add b.ResetTimer() here (after any setup, before the loop)
	// TODO: add the b.Loop() body: ConcatFast("ab", 100)
}
