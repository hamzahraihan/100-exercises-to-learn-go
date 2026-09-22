// Package filter teaches query parameters.
package filter

import "net/http"

// StatusParam returns the "status" query value, defaulting to "all".
func StatusParam(r *http.Request) string {
	if v := r.URL.Query().Get("status"); v != "" {
		return v
	}
	return "all"
}
