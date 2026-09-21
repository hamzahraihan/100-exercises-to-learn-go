// Package counter teaches sync.Mutex for shared state.
package counter

import "sync"

// Counter is a goroutine-safe int.
type Counter struct {
	mu sync.Mutex
	n  int
}

// Add increments by n.
func (c *Counter) Add(n int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n += n
}

// Value returns the current count.
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}
