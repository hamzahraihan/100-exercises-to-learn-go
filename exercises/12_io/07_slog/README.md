# Structured Logging with slog

`fmt.Fprintf(os.Stderr, "error: %v\n", err)` is honest but unstructured — a log pipeline can't parse it, filter by severity, or extract fields. `log/slog` (Go 1.21+) is the stdlib answer: structured, levelled, with JSON output built in.

## Concept 1 — The three-layer model

```go
handler := slog.NewJSONHandler(w, &slog.HandlerOptions{
    Level: slog.LevelInfo,
})
logger := slog.New(handler)
```

`Handler` decides where and how to write (JSON to `w`, minimum level). `Logger` is the API you call. They compose: swap `NewJSONHandler` for `NewTextHandler` and the logger API is identical.

## Concept 2 — Logging with key-value attrs

```go
logger.Info("ticket created",
    "id",     42,
    "title",  "Fix login",
    "status", "open",
)
// {"time":"...","level":"INFO","msg":"ticket created","id":42,"title":"Fix login","status":"open"}
```

Attributes follow the message as alternating key-value pairs. The handler encodes them; the caller never formats a string. A log aggregator can now filter `level=INFO` or search `status=open` across millions of lines.

## Concept 3 — Levels and filtering

```go
logger.Debug("verbose detail") // suppressed if Level >= Info
logger.Info("normal flow")     // emitted
logger.Warn("suspicious")
logger.Error("fatal path", "err", err)
```

`HandlerOptions.Level` is the gate. Debug is filtered at Info level — no allocation, no formatting, no write. Level filtering is cheap enough to leave calls in production code with the level turned up.

## Concept 4 — Writing to an io.Writer, not os.Stderr

Test with a `bytes.Buffer` to capture output without touching stderr:

```go
var buf bytes.Buffer
logger := NewLogger(&buf, slog.LevelInfo)
logger.Info("hello")
// inspect buf.String()
```

This is the testability contract: functions that accept `io.Writer` are testable; functions that write to a global logger are not.

<details><summary>Hint</summary>

Create the handler with `slog.NewJSONHandler(w, &slog.HandlerOptions{Level: lvl})`, then wrap it with `slog.New(handler)`. For `LogTicket`, pass alternating key-value pairs after the message: `logger.Info("ticket", "id", id, "title", title, "status", status)`.

</details>

## Task

Complete `NewLogger(w io.Writer, lvl slog.Level) *slog.Logger` in `logger.go` — return a `*slog.Logger` that writes JSON to `w` at the given minimum level.

Also complete `LogTicket(logger *slog.Logger, id int, title, status string)` — log one `Info` record with keys `"id"`, `"title"`, `"status"`.

## Check

```bash
go test ./exercises/12_io/07_slog/ -v
```
