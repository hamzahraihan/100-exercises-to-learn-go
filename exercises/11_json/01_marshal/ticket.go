// Package ticket teaches JSON marshaling.
package ticket

// Ticket is a support request.
type Ticket struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status"`
}

// MarshalTicket encodes t as JSON.
// TODO: use json.Marshal and return string(out) (import encoding/json).
func MarshalTicket(t Ticket) (string, error) {
	return "", nil
}
