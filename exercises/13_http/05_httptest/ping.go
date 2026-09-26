// Package ping teaches the httptest request/response cycle.
package ping

import "net/http"

// Ping answers liveness probes.
// TODO: write "pong" to w.
func Ping(w http.ResponseWriter, r *http.Request) {
}
