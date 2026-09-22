// Package hello teaches http.HandlerFunc basics.
package hello

import "net/http"

// Hello writes "hello". Unused params are fine; the imports must be used.
// TODO: write the body (hint: Fprintf from fmt, or io.WriteString).
func Hello(w http.ResponseWriter, r *http.Request) {
}
