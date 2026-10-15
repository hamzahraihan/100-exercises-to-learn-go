// Package validate teaches custom error types and errors.As.
package validate

import "fmt"

// Ticket is a support request.
type Ticket struct {
	Title       string
	Description string
}

// ValidationError carries the field that failed validation.
// TODO: implement Error() string to return "<Field>: <Message>".
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return "validation error" // TODO: format Field and Message
}

// NewTicket builds a Ticket, returning a wrapped *ValidationError
// when title or description is empty.
// TODO: check title (field="title") then description (field="description");
// wrap the *ValidationError with fmt.Errorf("new ticket: %w", &ValidationError{...}).
func NewTicket(title, description string) (Ticket, error) {
	return Ticket{title, description}, nil // TODO: validate
}

// ExtractField unwraps err looking for a *ValidationError and returns its Field.
// TODO: errors.As(err, &ve) to extract the concrete type.
func ExtractField(err error) (string, bool) {
	_ = fmt.Sprintf // keep import
	return "", false // TODO: implement
}
