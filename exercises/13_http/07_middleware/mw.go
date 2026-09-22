// Package mw teaches middleware.
package mw

import "net/http"

// WithHeader tags every response with X-Course: go, then calls next.
// TODO: w.Header().Set(...) before next.ServeHTTP(w, r).
func WithHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
