// Package shutdown teaches graceful shutdown.
package shutdown

import (
	"context"
	"net/http"
	"time"
)

// ShutdownGracefully stops srv, waiting up to timeout.
func ShutdownGracefully(srv *http.Server, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return srv.Shutdown(ctx)
}
