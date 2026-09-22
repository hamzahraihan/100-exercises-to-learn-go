// Package loadj teaches graceful corrupt-file handling.
package loadj

// Ticket is a stored support request.
type Ticket struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// LoadTickets decodes path; failures name the path (never panic).
// TODO: os.ReadFile + json.Unmarshal, wrap errors with path (import encoding/json, fmt, os).
func LoadTickets(path string) ([]Ticket, error) {
	return nil, nil
}
