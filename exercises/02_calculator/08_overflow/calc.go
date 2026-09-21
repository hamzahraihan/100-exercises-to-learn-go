package calc

// AddUint8 adds a and b, reporting wraparound. Go wraps silently,
// so callers must check: overflow is true when the result wrapped.
func AddUint8(a, b uint8) (sum uint8, overflow bool) {
	sum = a + b
	overflow = sum < a
	return sum, overflow
}
