// Package mw teaches middleware.
package mw

import "net/http"

// WithHeader tags every response with X-Course: go, then calls next.
func WithHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Course", "go")
		next.ServeHTTP(w, r)
	})
}
