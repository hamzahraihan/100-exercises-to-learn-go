// Package keys teaches the comparable constraint.
package keys

// Keys returns all map keys. K must be comparable (usable as a map key).
// TODO: range over m and append each k.
func Keys[K comparable, V any](m map[K]V) []K {
	return nil
}
