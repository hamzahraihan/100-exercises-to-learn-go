package fib

import "testing"

func TestFib(t *testing.T) {
	if got, want := Fib(10), 55; got != want {
		t.Fatalf("Fib(10) = %d, want %d", got, want)
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(20)
	}
}

func TestBenchmarkRuns(t *testing.T) {
	if res := testing.Benchmark(BenchmarkFib); res.N <= 0 {
		t.Fatal("benchmark did not run")
	}
}
