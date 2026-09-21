// Package must teaches converting panics to errors with recover.
package must

func parseOrPanic(s string) int {
	if s == "" {
		panic("empty input")
	}
	return len(s)
}

// MustParse returns len(s), or -1 when s is empty. No panic escapes.
// TODO: defer a func that recovers and sets the named return n to -1.
func MustParse(s string) (n int) {
	return parseOrPanic(s)
}
