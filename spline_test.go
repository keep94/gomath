package gomath_test

import (
	"testing"

	"github.com/keep94/gomath"
)

func TestZeroSpline(t *testing.T) {
	var s gomath.Spline
	assertPanic(t, func() { s.Eval(0.0) })
	assertPanic(t, func() { s.MinX() })
	assertPanic(t, func() { s.MaxX() })
}

func TestBadSpline(t *testing.T) {
	assertPanic(t, func() { gomath.NewSpline(nil) })
	assertPanic(t, func() { gomath.NewSpline([]gomath.Point{{0.0, 0.0}}) })
	assertPanic(t, func() {
		gomath.NewSpline([]gomath.Point{{5.0, 3.0}, {5.0, 6.0}})
	})
	assertPanic(t, func() {
		gomath.NewSpline([]gomath.Point{{5.0, 3.0}, {6.0, 6.0}, {5.0, 9.0}})
	})
}

func Test2Spline(t *testing.T) {
	s := gomath.NewSpline([]gomath.Point{{5.0, 10.0}, {7.0, 20.0}})
	assertCloseTo(t, 15.0, s.Eval(6.0))
	assertCloseTo(t, 20.0, s.Eval(7.0))
	assertCloseTo(t, 10.0, s.Eval(5.0))
	assertEqual(t, 5.0, s.MinX())
	assertEqual(t, 7.0, s.MaxX())
	assertPanic(t, func() { s.Eval(4.9) })
	assertPanic(t, func() { s.Eval(7.1) })
}

func Test4Spline(t *testing.T) {
	s := gomath.NewSpline(
		[]gomath.Point{{5.0, 10.0}, {7.0, 20.0}, {8.0, 40.0}, {11.0, 50.0}})
	assertCloseTo(t, 10.0, s.Eval(5.0))
	assertCloseTo(t, 20.0, s.Eval(7.0))
	assertCloseTo(t, 40.0, s.Eval(8.0))
	assertCloseTo(t, 50.0, s.Eval(11.0))
	assertTwiceDiff(t, s, 7.0, 0.1)
	assertTwiceDiff(t, s, 8.0, 0.1)
	assertCloseTo(t, 0.0, secondDerivative(s, 5.0, 0.1))
	assertCloseTo(t, 0.0, secondDerivative(s, 11.0, -0.1))
	assertEqual(t, 5.0, s.MinX())
	assertEqual(t, 11.0, s.MaxX())
	assertPanic(t, func() { s.Eval(4.9) })
	assertPanic(t, func() { s.Eval(11.1) })
}

func TestSplineWithSlopes(t *testing.T) {
	s := gomath.NewSplineWithSlopes(
		[]gomath.Point{{0.0, 0.0}, {1.0, 1.0}}, 0.0, 0.0)
	assertCloseTo(t, 0.0, s.Eval(0.0))
	assertCloseTo(t, 0.028, s.Eval(0.1))
	assertCloseTo(t, 0.104, s.Eval(0.2))
	assertCloseTo(t, 0.216, s.Eval(0.3))
	assertCloseTo(t, 0.352, s.Eval(0.4))
	assertCloseTo(t, 0.5, s.Eval(0.5))
	assertCloseTo(t, 0.648, s.Eval(0.6))
	assertCloseTo(t, 0.784, s.Eval(0.7))
	assertCloseTo(t, 0.896, s.Eval(0.8))
	assertCloseTo(t, 0.972, s.Eval(0.9))
	assertCloseTo(t, 1.0, s.Eval(1.0))

	s = gomath.NewSplineWithSlopes(
		[]gomath.Point{{0.0, 0.0}, {1.0, 1.0}}, 0.0, 2.0)
	assertCloseTo(t, 0.0, s.Eval(0.0))
	assertCloseTo(t, 0.01, s.Eval(0.1))
	assertCloseTo(t, 0.04, s.Eval(0.2))
	assertCloseTo(t, 0.09, s.Eval(0.3))
	assertCloseTo(t, 0.16, s.Eval(0.4))
	assertCloseTo(t, 0.25, s.Eval(0.5))
	assertCloseTo(t, 0.36, s.Eval(0.6))
	assertCloseTo(t, 0.49, s.Eval(0.7))
	assertCloseTo(t, 0.64, s.Eval(0.8))
	assertCloseTo(t, 0.81, s.Eval(0.9))
	assertCloseTo(t, 1.0, s.Eval(1.0))
}

func firstDerivative(s *gomath.Spline, x, diff float64) float64 {
	return (-11.0/6.0*s.Eval(x) + 3.0*s.Eval(x+diff) - 1.5*s.Eval(x+2*diff) + 1.0/3.0*s.Eval(x+3*diff)) / diff
}

func secondDerivative(s *gomath.Spline, x, diff float64) float64 {
	return (2.0*s.Eval(x) - 5.0*s.Eval(x+diff) + 4.0*s.Eval(x+2*diff) - s.Eval(x+3*diff)) / (diff * diff)
}

func assertTwiceDiff(t *testing.T, s *gomath.Spline, x, diff float64) {
	t.Helper()

	// First derivative to the left and right of x must match
	assertCloseTo(t, firstDerivative(s, x, -diff), firstDerivative(s, x, diff))

	// Second derivative to the left and right of x must match
	assertCloseTo(t, secondDerivative(s, x, -diff), secondDerivative(s, x, diff))
}
