// Package collect teaches buffered channels for fan-out collection.
package collect

// Collect returns [0, 2, ..., 2*(n-1)] computed by n goroutines sending
// over a buffered channel, collected in order.
// TODO: make(chan int, n), one goroutine per i sending 2*i, collect all.
func Collect(n int) []int {
	return make([]int, n)
}
