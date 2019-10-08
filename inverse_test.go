package gomath_test

import (
  "testing"

  "github.com/keep94/gomath"
)

func TestInverse(t *testing.T) {
  f := func(x float64) float64 { return x*x }
  assertCloseTo(t, 1.4142135, gomath.Inverse(f, 2.0, 1.0, 2.0))
  assertCloseTo(t, 1.7320508, gomath.Inverse(f, 3.0, 1.0, 2.0))
  assertCloseTo(t, 2.0, gomath.Inverse(f, 4.0, 1.0, 2.0))
  assertCloseTo(t, 2.0, gomath.Inverse(f, 5.0, 1.0, 2.0))
  assertCloseTo(t, 1.0, gomath.Inverse(f, 0.0, 1.0, 2.0))

  assertCloseTo(t, 1.4142135, gomath.Inverse(f, 2.0, 2.0, 1.0))
  assertCloseTo(t, 1.7320508, gomath.Inverse(f, 3.0, 2.0, 1.0))
  assertCloseTo(t, 2.0, gomath.Inverse(f, 4.0, 2.0, 1.0))
  assertCloseTo(t, 2.0, gomath.Inverse(f, 5.0, 2.0, 1.0))
  assertCloseTo(t, 1.0, gomath.Inverse(f, 0.0, 2.0, 1.0))

}

func TestInverseDecreasing(t *testing.T) {
  f := func(x float64) float64 { return 25.0 - x*x }
  assertCloseTo(t, 3.0, gomath.Inverse(f, 16.0, 1.1, 5.0))
  assertCloseTo(t, 4.0, gomath.Inverse(f, 9.0, 1.1, 5.0))
  assertCloseTo(t, 2.2360679, gomath.Inverse(f, 20.0, 1.1, 5.0))
  assertCloseTo(t, 1.1, gomath.Inverse(f, 23.79, 1.1, 5.0))
  assertCloseTo(t, 1.1, gomath.Inverse(f, 25.0, 1.1, 5.0))
  assertCloseTo(t, 5.0, gomath.Inverse(f, 0.0, 1.1, 5.0))
  assertCloseTo(t, 5.0, gomath.Inverse(f, -10.0, 1.1, 5.0))

  assertCloseTo(t, 3.0, gomath.Inverse(f, 16.0, 5.0, 1.1))
  assertCloseTo(t, 4.0, gomath.Inverse(f, 9.0, 5.0, 1.1))
  assertCloseTo(t, 2.2360679, gomath.Inverse(f, 20.0, 5.0, 1.1))
  assertCloseTo(t, 1.1, gomath.Inverse(f, 23.79, 5.0, 1.1))
  assertCloseTo(t, 1.1, gomath.Inverse(f, 25.0, 5.0, 1.1))
  assertCloseTo(t, 5.0, gomath.Inverse(f, 0.0, 5.0, 1.1))
  assertCloseTo(t, 5.0, gomath.Inverse(f, -10.0, 5.0, 1.1))

  assertCloseTo(t, 3.0, gomath.Inverse(f, 16.0, 3.0, 3.0))
  assertCloseTo(t, 3.0, gomath.Inverse(f, 17.0, 3.0, 3.0))
  assertCloseTo(t, 3.0, gomath.Inverse(f, 15.0, 3.0, 3.0))
}
