// Package must teaches converting panics to errors with recover.
package must

func parseOrPanic(s string) int {
	if s == "" {
		panic("empty input")
	}
	return len(s)
}

// MustParse returns len(s), or -1 when s is empty. No panic escapes.
func MustParse(s string) (n int) {
	defer func() {
		if recover() != nil {
			n = -1
		}
	}()
	return parseOrPanic(s)
}
