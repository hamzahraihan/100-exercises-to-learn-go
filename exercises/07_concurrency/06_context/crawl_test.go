package crawl

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestCrawlAll(t *testing.T) {
	urls := []string{"http://a.example", "http://b.example", "http://c.example"}
	fetch := func(u string) string { return u + "_ok" }

	results := make(chan string, len(urls))
	Crawl(context.Background(), urls, fetch, results)

	got := make(map[string]bool)
	for r := range results {
		got[r] = true
	}

	for _, u := range urls {
		want := u + "_ok"
		if !got[want] {
			t.Fatalf("got results %v, want %q in results (hint: send fetch(url) to results for every url)", got, want)
		}
	}
}

func TestCrawlCancelled(t *testing.T) {
	urls := []string{"http://1.example", "http://2.example", "http://3.example", "http://4.example", "http://5.example"}

	var calls atomic.Int64
	fetch := func(u string) string {
		calls.Add(1)
		return u + "_ok"
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel before Crawl is called

	results := make(chan string, 10)
	// Run in goroutine so we can observe the channel after a brief wait.
	go Crawl(ctx, urls, fetch, results)

	time.Sleep(50 * time.Millisecond)

	n := int64(len(results))
	c := calls.Load()
	if n != 0 || c != 0 {
		t.Fatalf("got %d results and %d fetch calls after pre-cancelled ctx, want 0 and 0 (hint: select on ctx.Done() at the top of each iteration and return immediately)", n, c)
	}
}
