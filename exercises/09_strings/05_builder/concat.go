// Package concat teaches strings.Builder.
package concat

import "strings"

// ConcatN repeats s n times efficiently.
func ConcatN(s string, n int) string {
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteString(s)
	}
	return b.String()
}
