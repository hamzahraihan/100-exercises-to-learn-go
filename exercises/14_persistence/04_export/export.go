// Package exporttk teaches JSON backup export.
package exporttk

// Ticket is a stored support request.
type Ticket struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// Export writes ts as a JSON array to path.
// TODO: json.Marshal + os.WriteFile(path, out, 0o644) (import encoding/json, os).
func Export(path string, ts []Ticket) error {
	return nil
}
