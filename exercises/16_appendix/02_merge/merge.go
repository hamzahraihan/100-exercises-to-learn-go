// Package merge teaches ordered merging by key.
package merge

// Ticket is a mergeable item.
type Ticket struct {
	ID    int
	Title string
}

// Merge overlays b onto a: same-ID tickets come from b, order follows a
// with b-only tickets appended in b's order.
func Merge(a, b []Ticket) []Ticket {
	pos := make(map[int]int, len(a))
	out := make([]Ticket, 0, len(a)+len(b))
	for _, t := range a {
		pos[t.ID] = len(out)
		out = append(out, t)
	}
	for _, t := range b {
		if i, ok := pos[t.ID]; ok {
			out[i] = t
		} else {
			pos[t.ID] = len(out)
			out = append(out, t)
		}
	}
	return out
}
