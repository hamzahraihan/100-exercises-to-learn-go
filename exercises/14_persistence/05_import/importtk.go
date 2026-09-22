// Package importtk teaches validated JSON import.
package importtk

// Ticket is a stored support request.
type Ticket struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// Import reads a JSON backup, validating every title.
// TODO: os.ReadFile + json.Unmarshal, reject empty titles (import encoding/json, errors, os).
func Import(path string) ([]Ticket, error) {
	return nil, nil
}
