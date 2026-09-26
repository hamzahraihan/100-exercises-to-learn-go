// Package status teaches custom JSON decoding.
package status

import "errors"

// Status is the ticket state.
type Status int

const (
	StatusOpen Status = iota
	StatusClosed
)

// Ticket carries a typed status.
type Ticket struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status Status `json:"status"`
}

// UnmarshalJSON decodes "open"/"closed", rejecting anything else.
// TODO: unquote data, switch on the word, assign through *s or error out.
func (s *Status) UnmarshalJSON(data []byte) error {
	return errors.New("TODO")
}
