# Context and Cancellation

Goroutines are cheap to start and expensive to forget. Every goroutine that
blocks on I/O, a channel, or a sleep without a stop signal is a goroutine you
cannot reclaim. `context.Context` is Go's first-class cancellation token —
passed explicitly, checked at every blocking point.

## The context tree

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel() // always defer — releases resources even if timeout fires first
```

`Background()` is the root; `WithTimeout`/`WithCancel`/`WithDeadline` derive
children. Cancel a parent and every child is cancelled. Defer cancel
unconditionally — leaving it out leaks a timer goroutine for the full timeout
duration even if the work finishes early.

## The select pattern for a cancellable worker

```go
func Worker(ctx context.Context, jobs <-chan int, results chan<- int) {
    for {
        select {
        case <-ctx.Done():
            return // context cancelled or timed out
        case j, ok := <-jobs:
            if !ok {
                return // channel closed
            }
            results <- j * 2
        }
    }
}
```

`ctx.Done()` returns a channel that closes when the context is cancelled. The
`select` races it against the jobs channel — whichever is ready wins. This is
how every well-behaved Go service stops cleanly: it doesn't poll a boolean or
check a global flag; it listens on a channel, just like any other piece of
concurrent Go code.

## Context travels as the first argument

Context is passed as the first argument to every function that does I/O,
waits, or starts goroutines. It is not global; it flows down the call tree.
The stdlib encodes this convention throughout: `http.Request.Context()`,
`sql.QueryContext`, `os.ReadFileContext`. If you find yourself reaching for a
package-level cancel bool, reach for a `context.Context` instead.

## Inspecting ctx.Err() after Done

```go
if err := ctx.Err(); err == context.DeadlineExceeded {
    // timed out
} else if err == context.Canceled {
    // explicitly cancelled
}
```

`ctx.Err()` returns nil until the context is done, then returns one of those
two sentinel values. You can log the reason, record a metric, or propagate it
as a wrapped error — all after the select that stopped the loop.

<details>
<summary>Hint</summary>

Inside the `for range urls` loop, use a `select` with two cases: `case <-ctx.Done(): return` (stop immediately), and `case results <- fetch(url):` (do the work and send the result). After the loop — or after an early return — call `close(results)` to signal the receiver that no more values are coming. Use `defer close(results)` at the top of the function so it fires on every return path.

</details>

## Task

Complete `Crawl` in `crawl.go`. The function receives a context, a slice of
URLs, a fetch function, and a results channel. It must call `fetch(url)` for
each URL and send the result to `results`, but stop immediately when `ctx` is
cancelled without processing the remaining URLs. Close `results` before
returning so the caller's range loop terminates.

```go
func Crawl(ctx context.Context, urls []string, fetch func(string) string, results chan<- string)
```

## Check

```bash
go test -race ./exercises/07_concurrency/06_context/ -v
```
