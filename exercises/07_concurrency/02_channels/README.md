# Channel Pipeline

A channel lets goroutines hand values back to the code that launched them:
each worker computes one result and sends it, and the caller collects the
sends into the output slice. The trick is keeping input order — send each
result tagged with its index (or write directly to `out[i]` from the
worker) and wait for every goroutine before returning. That is the Go
equivalent of the message-passing step in the Rust course this section is
adapted from.

## Task

Complete `DoubleAll` in `pipeline.go` so each input is doubled via a
goroutine and a channel, results land in input order, and the test passes.

## Check

```bash
go test ./exercises/07_concurrency/02_channels/ -v
```
