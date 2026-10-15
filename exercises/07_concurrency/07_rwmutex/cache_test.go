package rwcache

import (
	"fmt"
	"sync"
	"testing"
)

func TestCacheBasic(t *testing.T) {
	c := NewCache()

	// Miss on empty cache.
	if _, ok := c.Get("missing"); ok {
		t.Fatalf("got ok=true for missing key, want false")
	}

	// Set then Get.
	c.Set("lang", "go")
	v, ok := c.Get("lang")
	if !ok {
		t.Fatalf("got ok=false after Set, want true")
	}
	if v != "go" {
		t.Fatalf("got %v, want %v (hint: Set should store the value so Get can retrieve it)", v, "go")
	}

	// Overwrite.
	c.Set("lang", "Go")
	v, ok = c.Get("lang")
	if !ok {
		t.Fatalf("got ok=false after overwrite, want true")
	}
	if v != "Go" {
		t.Fatalf("got %v, want %v (hint: Set should overwrite the existing value for the same key)", v, "Go")
	}

	// Second miss.
	if _, ok := c.Get("other"); ok {
		t.Fatalf("got ok=true for key \"other\" which was never set, want false")
	}
}

func TestCacheConcurrent(t *testing.T) {
	c := NewCache()

	const (
		numWriters    = 5
		keysPerWriter = 10
		numReaders    = 20
		keysPerReader = 10
		totalKeys     = numWriters * keysPerWriter // 50
	)

	// Pre-populate so readers have keys to hit from the start.
	for i := range totalKeys {
		c.Set(fmt.Sprintf("key-%d", i), fmt.Sprintf("init-%d", i))
	}

	var wg sync.WaitGroup

	// Writers: each overwrites its slice of keys.
	for w := range numWriters {
		w := w
		wg.Add(1)
		go func() {
			defer wg.Done()
			base := w * keysPerWriter
			for i := range keysPerWriter {
				k := fmt.Sprintf("key-%d", base+i)
				c.Set(k, fmt.Sprintf("writer-%d-val-%d", w, i))
			}
		}()
	}

	// Readers: each reads a spread of keys.
	for r := range numReaders {
		r := r
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range keysPerReader {
				k := fmt.Sprintf("key-%d", (r*keysPerReader+i)%totalKeys)
				c.Get(k) // value may be init or writer-updated; either is fine
			}
		}()
	}

	wg.Wait()

	// Spot-check: every key must still be readable (may hold init or writer value).
	for i := range totalKeys {
		k := fmt.Sprintf("key-%d", i)
		if _, ok := c.Get(k); !ok {
			t.Fatalf("got ok=false for %q after concurrent writes, want true (hint: check your mutex protects the map correctly)", k)
		}
	}
}
