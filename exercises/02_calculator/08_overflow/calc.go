package calc

// AddUint8 adds a and b, reporting wraparound. Go wraps silently,
// so callers must check: overflow is true when the result wrapped.
// TODO: compute sum and set overflow = sum < a.
func AddUint8(a, b uint8) (sum uint8, overflow bool) {
	return a + b, false
}
