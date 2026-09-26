// Package merge teaches ordered merging by key.
package merge

// Ticket is a mergeable item.
type Ticket struct {
	ID    int
	Title string
}

// Merge overlays b onto a: same-ID tickets come from b, order follows a
// with b-only tickets appended in b's order.
// TODO: index a by ID, walk b overwriting known ids and appending newcomers.
func Merge(a, b []Ticket) []Ticket {
	return nil
}
