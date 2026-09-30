package shutdown

import (
	"net"
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

func TestShutdownGracefullyStopsListener(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	srv := &http.Server{Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})}
	served := make(chan error, 1)
	go func() { served <- srv.Serve(ln) }()
	// Give Serve a moment to start; deadline keeps this deterministic.
	deadline := time.Now().Add(2 * time.Second)
	for {
		conn, err := net.DialTimeout("tcp", ln.Addr().String(), 100*time.Millisecond)
		if err == nil {
			conn.Close()
			break
		}
		if time.Now().After(deadline) {
			t.Fatalf("server did not start: %v", err)
		}
		time.Sleep(10 * time.Millisecond)
	}

	start := time.Now()
	if err := ShutdownGracefully(srv, 2*time.Second); err != nil {
		t.Fatalf("ShutdownGracefully(running) = %v, want nil", err)
	}
	if elapsed := time.Since(start); elapsed > 2*time.Second {
		t.Fatalf("shutdown took %v, want < 2s", elapsed)
	}
	if err := <-served; err != http.ErrServerClosed {
		t.Fatalf("Serve = %v, want ErrServerClosed", err)
	}
}
