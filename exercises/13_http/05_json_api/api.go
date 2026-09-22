// Package api teaches JSON request/response handling.
package api

import "net/http"

// Ticket is a support request.
type Ticket struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status"`
}

// CreateTicket decodes a ticket, validates the title, and echoes it back
// with 201 as JSON (400 for bad input).
// TODO: json.Decode the body, check Title, set Content-Type, WriteHeader,
// json.Encode the ticket (import encoding/json).
func CreateTicket(w http.ResponseWriter, r *http.Request) {
}
