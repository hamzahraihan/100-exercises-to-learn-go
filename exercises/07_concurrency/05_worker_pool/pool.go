// Package pool teaches the worker-pool concurrency pattern.
package pool

// SumSquares returns the sum of squares of all inputs using a fixed
// pool of worker goroutines.
// TODO: create a jobs channel (buffered, cap=len(inputs)) and a results
// channel (buffered, cap=len(inputs)), launch `workers` goroutines each
// ranging over jobs and sending i*i to results, fill jobs then close it,
// drain results and sum them.
func SumSquares(inputs []int, workers int) int {
	var sum int
	for _, v := range inputs {
		sum += v * v // TODO: replace with worker pool
	}
	return sum // stub: correct answer, wrong mechanism — must use goroutines
}
