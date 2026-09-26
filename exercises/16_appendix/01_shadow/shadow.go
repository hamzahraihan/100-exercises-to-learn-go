// Package shadow teaches variable shadowing hazards.
package shadow

// ApplyBonus adds bonus to total when bonus is positive, using assignment
// so the parameter itself is updated rather than shadowed.
func ApplyBonus(total, bonus int) int {
	if bonus > 0 {
		total += bonus
	}
	return total
}
