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
func Page(ts []Ticket, offset, limit int) []Ticket {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 || offset >= len(ts) {
		return []Ticket{}
	}
	end := offset + limit
	if end > len(ts) {
		end = len(ts)
	}
	return ts[offset:end]
}
