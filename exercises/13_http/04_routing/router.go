// Package router teaches ServeMux method+path patterns (Go 1.22+).
package router

import "net/http"

// NewMux routes the ticket collection and single-ticket reads.
// TODO: mux.HandleFunc("GET /tickets", ...) writing "list" and
// "GET /tickets/{id}" writing "one:"+r.PathValue("id").
func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
