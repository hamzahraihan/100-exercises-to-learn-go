// Package pipeline teaches goroutines + channels.
package pipeline

// DoubleAll doubles each number using a goroutine and channel,
// preserving input order.
func DoubleAll(nums []int) []int {
	out := make([]int, len(nums))
	type result struct {
		idx int
		val int
	}
	ch := make(chan result, len(nums))
	for i, n := range nums {
		go func(i, n int) {
			ch <- result{idx: i, val: n * 2}
		}(i, n)
	}
	for range nums {
		r := <-ch
		out[r.idx] = r.val
	}
	return out
}
