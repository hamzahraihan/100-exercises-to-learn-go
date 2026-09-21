// Package counter teaches sync.Mutex for shared state.
package counter

import "sync"

// Counter is a goroutine-safe int.
// TODO: add a sync.Mutex field and guard n.
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
