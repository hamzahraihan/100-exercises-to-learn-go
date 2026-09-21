package calc

// ToInt64 converts x explicitly. Go has no implicit numeric
// conversions (unlike some languages) and no `as` casts:
// every conversion is written T(x).
func ToInt64(x int32) int64 {
	return int64(x)
}
