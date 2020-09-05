package gomath

// Inverse returns x such that f(x) = y.
// Caller must choose lower and upper so that x falls in between them.
// f must be monotone increasing or decreasing between lower and upper.
func Inverse(
	f func(float64) float64,
	y float64,
	lower float64,
	upper float64) float64 {
	var g func(float64) float64
	if f(lower) > f(upper) {
		g = func(val float64) float64 {
			return y - f(val)
		}
	} else {
		g = func(val float64) float64 {
			return f(val) - y
		}
	}
	step := (upper - lower) / 2.0
	result := (upper + lower) / 2.0
	for {
		step /= 2.0
		var nextResult float64
		if g(result) > 0.0 {
			nextResult = result - step
		} else {
			nextResult = result + step
		}
		if nextResult == result {
			return result
		}
		result = nextResult
	}
}
