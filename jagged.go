package gomath

import (
	"math"
)

// Jagged is a function that is continuous everywhere but differentiable
// nowhere.
func Jagged(x float64) float64 {
	x = x - math.Floor(x)
	mult := 1.0
	result := 0.0
	oldResult := 0.0
	for {
		if x > 0.5 {
			result += (1.0 - x) * mult
		} else {
			result += x * mult
		}
		if result == oldResult {
			return result
		}
		oldResult = result
		mult /= 2.0
		x *= 2.0
		if x >= 1.0 {
			x -= 1.0
		}
	}
	return 0.0
}
