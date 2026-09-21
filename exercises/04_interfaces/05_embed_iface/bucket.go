// Package bucket teaches interface embedding.
package bucket

// Reader reads stored data.
type Reader interface {
	Read() string
}

// Writer stores data.
type Writer interface {
	Write(s string)
}

// ReadWriter embeds both: one interface, two capabilities.
type ReadWriter interface {
	Reader
	Writer
}

// Bucket is an in-memory store.
type Bucket struct {
	data string
}

// Write stores s.
// TODO: assign b.data = s.
func (b *Bucket) Write(s string) {
}

// Read returns stored data.
// TODO: return b.data.
func (b *Bucket) Read() string {
	return ""
}

var _ ReadWriter = (*Bucket)(nil)
