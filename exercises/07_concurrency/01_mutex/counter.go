// Package counter teaches sync.Mutex for shared state.
package counter

import "sync"

// Counter is a goroutine-safe int.
// TODO: guard n with mu in Add and Value (Lock/Unlock).
type Counter struct {
	mu sync.Mutex
	n  int
}

// Add increments by n.
func (c *Counter) Add(n int) {
	c.n += n // TODO: Lock/Unlock
}

// Value returns the current count.
func (c *Counter) Value() int {
	return c.n // TODO: Lock/Unlock
}
