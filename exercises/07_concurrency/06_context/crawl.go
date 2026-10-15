// Package crawl teaches context.Context for goroutine cancellation.
package crawl

import "context"

// Crawl calls fetch(url) for each url and sends the result to results.
// It must stop and close results when ctx is cancelled.
// TODO: range over urls; in each iteration select on ctx.Done() (return)
// and a "do the work" case; close results before returning.
func Crawl(ctx context.Context, urls []string, fetch func(string) string, results chan<- string) {
	for _, url := range urls {
		results <- fetch(url) // TODO: check ctx.Done() before each fetch
	}
	close(results)
}
