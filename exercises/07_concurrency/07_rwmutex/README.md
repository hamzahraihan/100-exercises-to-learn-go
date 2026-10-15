# Read/Write Mutex

The mutex lesson put one lock on everything. That's safe, but heavy: reads can't race each other — only a read-write conflict is dangerous. When reads dominate (a cache hit ratio of 99%), locking out every concurrent reader is waste. `sync.RWMutex` splits the lock in two.

## The two locks

```go
var mu sync.RWMutex

// Readers: many can hold this at once
mu.RLock()
defer mu.RUnlock()
v := cache[key]

// Writers: exclusive; waits for all readers to finish
mu.Lock()
defer mu.Unlock()
cache[key] = value
```

`RLock` grants shared access; many goroutines may hold it simultaneously. `Lock` waits until every `RLock` is released, then blocks all new `RLock` calls. The invariant: a writer is never concurrent with anyone. A reader is never concurrent with a writer. Two readers are fine.

## When to reach for it

Rule of thumb: if reads outnumber writes by 10:1 or more, `RWMutex` pays off. Below that ratio the added complexity isn't justified — a plain `Mutex` is simpler and nearly as fast. Profile first; don't swap locks on instinct.

## The trap — don't upgrade a read lock

```go
// WRONG — deadlock
mu.RLock()
if _, ok := cache[key]; !ok {
    mu.Lock() // blocks forever: RLock is still held
    cache[key] = compute()
    mu.Unlock()
}
mu.RUnlock()
```

You cannot "upgrade" an `RLock` to a `Lock`. Release the read lock first, re-acquire the write lock, then check again (the key may have appeared in the gap).

> **Note on the stub:** The stub uses a plain `Mutex` — correct but it serialises all reads. Switch to `RWMutex` so concurrent reads proceed in parallel; `-race` proves the upgrade is safe.

<details><summary>Hint</summary>

Embed `sync.RWMutex` in the struct. In `Get`, call `mu.RLock()` / `defer mu.RUnlock()`. In `Set`, call `mu.Lock()` / `defer mu.Unlock()`. That's the entire change — the correctness story stays identical; only the locking granularity improves.

</details>

## Task

Complete `Get` and `Set` on the `Cache` type in `cache.go`. Replace the plain `sync.Mutex` with `sync.RWMutex`, use `RLock`/`RUnlock` in `Get` and `Lock`/`Unlock` in `Set`. The test hammers it with 20 concurrent readers and 5 concurrent writers; `-race` must pass clean.

## Check

```bash
go test -race ./exercises/07_concurrency/07_rwmutex/ -v
```
