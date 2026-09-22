// Package errw teaches structured error responses.
package errw

import (
	"encoding/json"
	"net/http"
)

// WriteError answers code with a {"error":msg} JSON body.
func WriteError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
