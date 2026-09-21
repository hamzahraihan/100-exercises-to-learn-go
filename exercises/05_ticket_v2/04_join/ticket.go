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
func ValidateAll(title, desc string) error {
	var errs []error
	if title == "" {
		errs = append(errs, ErrBadTitle)
	}
	if len(desc) < 10 {
		errs = append(errs, ErrBadDesc)
	}
	return errors.Join(errs...)
}
