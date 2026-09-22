// Package errw teaches structured error responses.
package errw

import "net/http"

// WriteError answers code with a {"error":msg} JSON body.
// TODO: header, WriteHeader(code), encode map/struct (import encoding/json).
func WriteError(w http.ResponseWriter, code int, msg string) {
}
