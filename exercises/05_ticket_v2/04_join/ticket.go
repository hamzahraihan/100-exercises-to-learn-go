// Package ticket teaches joining multiple errors.
package ticket

import "errors"

var (
	// ErrBadTitle is returned for empty titles.
	ErrBadTitle = errors.New("bad title")
	// ErrBadDesc is returned for short descriptions.
	ErrBadDesc = errors.New("bad description")
)

// ValidateAll checks title and description, reporting every problem.
// TODO: collect per-field errors and combine with errors.Join (nil when valid).
func ValidateAll(title, desc string) error {
	return errors.New("TODO")
}
