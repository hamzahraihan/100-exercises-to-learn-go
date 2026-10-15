# Worker Pool

Fan-out/fan-in is Go's workhorse concurrency pattern. Raw goroutine-per-item works
until it doesn't — 10 000 goroutines each doing I/O will crush the OS scheduler and
exhaust file descriptors long before they finish. A **worker pool** caps parallelism
at a fixed number while an unbounded queue of jobs keeps feeding them.

## The shape

```go
jobs    := make(chan int, len(inputs))
results := make(chan int, len(inputs))

for w := 0; w < workers; w++ {
    go func() {
        for j := range jobs {
            results <- j * j
        }
    }()
}

for _, v := range inputs {
    jobs <- v
}
close(jobs)

for range inputs {
    _ = <-results
}
```

Two channels carry all the work:

- `jobs` — the queue. The main goroutine fills it then **closes** it. `close` is the
  signal that tells every worker's `range` loop to exit cleanly once the queue is
  empty. Without the close, workers block forever waiting for a job that never comes.
- `results` — the output lane. Each worker sends its answer here; the main goroutine
  drains exactly `len(inputs)` values to collect all results.

## Why buffered?

`jobs` is buffered to `len(inputs)` so the main goroutine can fill the whole queue
before any worker starts — no blocking, no deadlock.

`results` is also buffered to `len(inputs)`. If it were unbounded, workers writing to
`results` would block when the collector hasn't caught up. Sizing it to the total
number of expected results means workers can always send immediately.

## Order caveat

Results arrive in **arbitrary order** — which worker wins the race to process a given
job is scheduling-dependent. When you need ordered results, tag each result with its
input index (the same trick shown in `02_channels`). This exercise collects into a
sum, so order doesn't matter; that keeps the focus on the pool structure itself.

<details>
<summary>Hint</summary>

1. `jobs := make(chan int, len(inputs))` — buffered queue.
2. `results := make(chan int, len(inputs))` — buffered collector.
3. Launch exactly `workers` goroutines; each should `for j := range jobs { results <- j*j }`.
4. Fill `jobs` with every value from `inputs`, then `close(jobs)`.
5. Loop `len(inputs)` times reading from `results` and accumulating the sum.

</details>

## Task

Complete `SumSquares(inputs []int, workers int) int` in `pool.go`.

The function must square each element of `inputs` using **exactly `workers`**
goroutines arranged as a worker pool, then return the sum of all squares.

> **Note:** The starter stub returns the correct answer the wrong way (a plain `for`
> loop). The tests will pass against the stub, but your implementation must actually
> use goroutines. Run with `-race` and inspect the code: a real pool will show
> concurrent goroutine access on the results channel under the race detector.

## Check

```bash
go test -race ./exercises/07_concurrency/05_worker_pool/ -v
```
