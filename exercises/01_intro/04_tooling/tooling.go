// Package tooling teaches gofmt and go vet.
package tooling

// Quote wraps s in double quotes.
func Quote(s string) string {
	return "\"" + s + "\""
}
