// Package kind teaches type switches.
package kind

// Describe names supported dynamic types.
// TODO: switch v := v.(type) for int and string cases.
func Describe(v any) string {
	return "unknown"
}
