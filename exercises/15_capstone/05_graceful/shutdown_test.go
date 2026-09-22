package shutdown

import (
	"net/http"
	"testing"
	"time"
)

func TestShutdownGracefully(t *testing.T) {
	srv := &http.Server{}
	if err := ShutdownGracefully(srv, 2*time.Second); err != nil {
		t.Fatalf("ShutdownGracefully = %v, want nil", err)
	}
}
