// Package router teaches ServeMux method+path patterns (Go 1.22+).
package router

import (
	"io"
	"net/http"
)

// NewMux routes the ticket collection and single-ticket reads.
func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /tickets", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "list")
	})
	mux.HandleFunc("GET /tickets/{id}", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "one:"+r.PathValue("id"))
	})
	return mux
}
