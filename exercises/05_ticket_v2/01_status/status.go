// Package status teaches iota enums (Go's answer to simple Rust enums).
package status

// Status is a ticket state.
type Status int

const (
	StatusOpen Status = iota
	StatusInProgress
	StatusClosed
)

// Valid reports whether s is a known status.
func (s Status) Valid() bool {
	switch s {
	case StatusOpen, StatusInProgress, StatusClosed:
		return true
	default:
		return false
	}
}
