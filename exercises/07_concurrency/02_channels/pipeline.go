// Package pipeline teaches goroutines + channels.
package pipeline

// DoubleAll doubles each number using a goroutine and channel,
// preserving input order.
// TODO: launch a goroutine per index (or one worker), send results
// over a channel, collect in order.
func DoubleAll(nums []int) []int {
	out := make([]int, len(nums))
	return out
}
