// Package collect teaches buffered channels for fan-out collection.
package collect

// Collect returns [0, 2, ..., 2*(n-1)] computed by n goroutines sending
// over a buffered channel, collected in order.
func Collect(n int) []int {
	type pair struct {
		i, v int
	}
	ch := make(chan pair, n)
	for i := 0; i < n; i++ {
		go func(i int) { ch <- pair{i, 2 * i} }(i)
	}
	out := make([]int, n)
	for k := 0; k < n; k++ {
		p := <-ch
		out[p.i] = p.v
	}
	return out
}
