// Package store teaches pagination over slices.
package store

// Ticket is a stored item.
type Ticket struct {
	ID     int
	Title  string
	Closed bool
}

// Page returns up to limit tickets starting at offset. Out-of-range or
// non-positive inputs yield an empty (len 0) result, never a panic.
// TODO: clamp offset/limit into range, then slice.
func Page(ts []Ticket, offset, limit int) []Ticket {
	return nil
}
