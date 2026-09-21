package counter

import (
	"sync"
	"testing"
)

func TestConcurrentAdds(t *testing.T) {
	var c Counter
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); c.Add(1) }()
	}
	wg.Wait()
	if got := c.Value(); got != 100 {
		t.Fatalf("Value() = %d, want 100 (hint: sync.Mutex)", got)
	}
}
