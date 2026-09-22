// Package hello teaches http.HandlerFunc basics.
package hello

import (
	"io"
	"net/http"
)

// Hello writes "hello". Unused params are fine; the imports must be used.
func Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello")
}
