package calc

// Perm is a set of permission flags.
type Perm uint8

const (
	Read Perm = 1 << iota
	Write
	Execute
)

// Has reports whether flag is set in p.
// TODO: return p&flag == flag.
func Has(p, flag Perm) bool {
	return false
}
