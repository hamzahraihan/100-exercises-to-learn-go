// Package store teaches sorting and filtering slices.
package store

// Ticket is a stored item.
type Ticket struct {
	ID     int
	Title  string
	Closed bool
}

// SortedByID returns tickets ordered by ID ascending.
// TODO: sort a copy (hint: slices.SortFunc or sort.Slice).
func SortedByID(ts []Ticket) []Ticket {
	return nil
}

// OpenOnly keeps tickets that are not closed.
// TODO: append matches into a fresh slice.
func OpenOnly(ts []Ticket) []Ticket {
	return nil
}
