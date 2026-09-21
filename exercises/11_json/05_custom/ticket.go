// Package ticket teaches custom JSON marshaling.
package ticket

import (
	"encoding/json"
	"errors"
	"strconv"
)

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
func (s Status) MarshalJSON() ([]byte, error) {
	switch s {
	case StatusOpen:
		return []byte(strconv.Quote("open")), nil
	case StatusClosed:
		return []byte(strconv.Quote("closed")), nil
	default:
		return nil, errors.New("unknown status")
	}
}

// MarshalTicket encodes t as JSON (status via MarshalJSON).
func MarshalTicket(t Ticket) (string, error) {
	out, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(out), nil
}
