// Package shutdown teaches graceful shutdown.
package shutdown

import (
	"errors"
	"net/http"
	"time"
)

// ShutdownGracefully stops srv, waiting up to timeout.
// TODO: context.WithTimeout + srv.Shutdown (import context in the fix).
func ShutdownGracefully(srv *http.Server, timeout time.Duration) error {
	return errors.New("TODO")
}
