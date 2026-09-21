// Package status teaches iota enums (Go's answer to simple Rust enums).
package status

// Status is a ticket state.
// TODO: define Status as int-based type with iota constants
// StatusOpen, StatusInProgress, StatusClosed, and a Valid() method.
type Status int

const (
	StatusOpen Status = iota
	StatusInProgress
	StatusClosed
)

// Valid reports whether s is a known status.
func (s Status) Valid() bool {
	return false // TODO: switch on s
}
