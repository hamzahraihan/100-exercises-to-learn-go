// Package ticket teaches custom JSON marshaling.
package ticket

import "errors"

// Status is the ticket state.
type Status int

const (
	StatusOpen Status = iota
	StatusClosed
)

// Ticket is a support request with a typed status.
type Ticket struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status Status `json:"status"`
}

// MarshalJSON renders the status as "open"/"closed".
// TODO: switch on s and return the quoted string (strconv.Quote helps).
func (s Status) MarshalJSON() ([]byte, error) {
	return nil, errors.New("TODO")
}

// MarshalTicket encodes t as JSON (status via MarshalJSON).
// TODO: use json.Marshal (import encoding/json).
func MarshalTicket(t Ticket) (string, error) {
	return "", nil
}
