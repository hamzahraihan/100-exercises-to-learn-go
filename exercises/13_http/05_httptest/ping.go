// Package ping teaches the httptest request/response cycle.
package ping

import (
	"io"
	"net/http"
)

// Ping answers liveness probes.
func Ping(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong")
}
