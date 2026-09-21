// Package ticket teaches JSON marshaling.
package ticket

import "encoding/json"

// Ticket is a support request.
type Ticket struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status"`
}

// MarshalTicket encodes t as JSON.
func MarshalTicket(t Ticket) (string, error) {
	out, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(out), nil
}
