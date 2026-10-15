package pool

import "testing"

func TestSumSquaresSmall(t *testing.T) {
	// 1² + 2² + 3² + 4² + 5² = 1 + 4 + 9 + 16 + 25 = 55
	got := SumSquares([]int{1, 2, 3, 4, 5}, 3)
	want := 55
	if got != want {
		t.Fatalf("got %d, want %d (hint: sum each input squared, not the inputs themselves)", got, want)
	}
}

func TestSumSquaresSingleWorker(t *testing.T) {
	// 10² + 20² + 30² = 100 + 400 + 900 = 1400
	got := SumSquares([]int{10, 20, 30}, 1)
	want := 1400
	if got != want {
		t.Fatalf("got %d, want %d (hint: works=1 still uses a goroutine; the pool just has one worker)", got, want)
	}
}

func TestSumSquaresLarger(t *testing.T) {
	// sum of i² for i in [1..100] = 100*101*201/6 = 338350
	inputs := make([]int, 100)
	for i := range inputs {
		inputs[i] = i + 1
	}
	got := SumSquares(inputs, 8)
	want := 338350
	if got != want {
		t.Fatalf("got %d, want %d (hint: make sure every job is sent to a worker and every result is collected)", got, want)
	}
}

func TestSumSquaresWorkerCountEqualsInputs(t *testing.T) {
	// workers == len(inputs): each worker handles exactly one job
	inputs := []int{3, 5, 7}
	got := SumSquares(inputs, len(inputs))
	want := 83 // 9 + 25 + 49
	if got != want {
		t.Fatalf("got %d, want %d (hint: workers may exceed active jobs — workers block on range and exit when jobs is closed)", got, want)
	}
}
