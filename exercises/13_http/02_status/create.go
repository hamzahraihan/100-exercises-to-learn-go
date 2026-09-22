// Package created teaches status codes.
package created

import "net/http"

// Create answers resource creation.
func Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
