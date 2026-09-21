// Package keys teaches the comparable constraint.
package keys

// Keys returns all map keys. K must be comparable (usable as a map key).
func Keys[K comparable, V any](m map[K]V) []K {
	ks := make([]K, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}
