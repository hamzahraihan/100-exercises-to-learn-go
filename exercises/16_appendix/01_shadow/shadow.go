// Package shadow teaches variable shadowing hazards.
package shadow

// ApplyBonus adds bonus to total when bonus is positive.
// TODO: add bonus to total inside the if. Beware: declaring
// `total := ...` there would shadow the parameter instead of updating it.
func ApplyBonus(total, bonus int) int {
	return total
}
