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
func (b *Bucket) Write(s string) {
	b.data = s
}

// Read returns stored data.
func (b *Bucket) Read() string {
	return b.data
}

var _ ReadWriter = (*Bucket)(nil)
