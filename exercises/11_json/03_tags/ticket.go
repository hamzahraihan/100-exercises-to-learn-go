// Package ticket teaches JSON struct tags.
package ticket

import "encoding/json"

// Ticket is a support request. InternalNote never leaves the process.
type Ticket struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description,omitempty"`
	Status       string `json:"status"`
	InternalNote string `json:"-"`
}

// MarshalTicket encodes t as JSON.
func MarshalTicket(t Ticket) (string, error) {
	out, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(out), nil
}
