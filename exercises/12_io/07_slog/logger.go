// Package slogex teaches log/slog for structured, levelled logging.
package slogex

import (
	"io"
	"log/slog"
)

// NewLogger returns a *slog.Logger writing JSON to w at the given minimum level.
// TODO: slog.NewJSONHandler(w, &slog.HandlerOptions{Level: lvl}), then slog.New(handler).
func NewLogger(w io.Writer, lvl slog.Level) *slog.Logger {
	return slog.New(slog.NewTextHandler(w, nil)) // TODO: JSON, respect lvl
}

// LogTicket logs one Info record with keys "id", "title", "status".
// TODO: logger.Info("ticket", "id", id, "title", title, "status", status).
func LogTicket(logger *slog.Logger, id int, title, status string) {
	logger.Info("ticket") // TODO: add id, title, status attrs
}
