// Package rwcache teaches sync.RWMutex for read-heavy shared state.
package rwcache

import "sync"

// Cache is a goroutine-safe string→string map.
// TODO: replace sync.Mutex with sync.RWMutex;
// use RLock/RUnlock in Get and Lock/Unlock in Set.
type Cache struct {
	mu sync.Mutex
	m  map[string]string
}

// NewCache returns an initialised Cache.
func NewCache() *Cache {
	return &Cache{m: make(map[string]string)}
}

// Get returns the value for key and whether it was found.
func (c *Cache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.m[key]
	return v, ok
}

// Set stores value under key.
func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = value
}
