// Package filter teaches query parameters.
package filter

import "net/http"

// StatusParam returns the "status" query value, defaulting to "all".
// TODO: r.URL.Query().Get with an empty check.
func StatusParam(r *http.Request) string {
	return ""
}
