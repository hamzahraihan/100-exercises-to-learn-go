// Package order teaches implicit interfaces via fmt.Stringer.
package order

// Order is a purchase.
type Order struct {
	ID   int
	Item string
}

// String implements fmt.Stringer. Go interfaces are implicit:
// no "implements" keyword (unlike Rust trait impls).
// TODO: return a string containing the id and item, e.g. fmt.Sprintf("Order %d: %s", o.ID, o.Item).
func (o Order) String() string {
	return ""
}
