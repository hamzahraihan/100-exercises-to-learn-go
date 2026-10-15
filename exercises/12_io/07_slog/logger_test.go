package slogex

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"testing"
)

// TestNewLoggerJSON verifies that NewLogger produces valid JSON output.
func TestNewLoggerJSON(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, slog.LevelInfo)
	logger.Info("hello")

	raw := buf.Bytes()
	if len(raw) == 0 {
		t.Fatal("got empty output, want JSON log line")
	}

	var m map[string]any
	if err := json.Unmarshal(raw, &m); err != nil {
		t.Fatalf("got %v, want valid JSON (hint: use slog.NewJSONHandler, not NewTextHandler)", err)
	}

	got, ok := m["level"]
	if !ok {
		t.Fatalf("got no \"level\" key, want \"INFO\" (hint: JSON handler encodes level as a string field)")
	}
	if got != "INFO" {
		t.Fatalf("got %v, want %v (hint: Info() should produce level INFO)", got, "INFO")
	}
}

// TestLogTicketFields verifies that LogTicket emits all required key-value attributes.
func TestLogTicketFields(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, slog.LevelInfo)
	LogTicket(logger, 7, "Bug", "open")

	var m map[string]any
	if err := json.Unmarshal(buf.Bytes(), &m); err != nil {
		t.Fatalf("got %v, want valid JSON (hint: NewLogger must use NewJSONHandler)", err)
	}

	if got, want := m["msg"], "ticket"; got != want {
		t.Fatalf("got %v, want %v (hint: first arg to logger.Info is the message)", got, want)
	}
	if got, want := m["id"], float64(7); got != want {
		t.Fatalf("got %v, want %v (hint: pass \"id\", id as key-value pair after message)", got, want)
	}
	if got, want := m["title"], "Bug"; got != want {
		t.Fatalf("got %v, want %v (hint: pass \"title\", title as key-value pair after message)", got, want)
	}
	if got, want := m["status"], "open"; got != want {
		t.Fatalf("got %v, want %v (hint: pass \"status\", status as key-value pair after message)", got, want)
	}
}

// TestLevelFiltering verifies that records below the minimum level are suppressed.
func TestLevelFiltering(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, slog.LevelWarn)

	logger.Info("ignored")
	if buf.Len() != 0 {
		t.Fatalf("got %v, want empty output (hint: Info is below Warn — HandlerOptions.Level must gate output)", buf.String())
	}

	logger.Warn("seen")
	if buf.Len() == 0 {
		t.Fatalf("got empty output, want non-empty (hint: Warn >= Warn should be emitted)")
	}
}
