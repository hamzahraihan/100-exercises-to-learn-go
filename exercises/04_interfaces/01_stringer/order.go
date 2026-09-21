// Package order teaches implicit interfaces via fmt.Stringer.
package order

import "fmt"

// Order is a purchase.
type Order struct {
	ID   int
	Item string
}

// String implements fmt.Stringer. Go interfaces are implicit:
// no "implements" keyword (unlike Rust trait impls).
func (o Order) String() string {
	return fmt.Sprintf("Order %d: %s", o.ID, o.Item)
}
