package gomath_test

import (
  "math/big"
  "testing"

  "github.com/keep94/gomath"
)

func TestUgly(t *testing.T) {
  uglies := gomath.Ugly(2, 3, 5)
  checkInfBigIntStream(
      t,
      uglies,
      1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18, 20, 24, 25, 27,
      30, 32, 36, 40, 45, 48, 50, 54, 60, 64, 72, 75, 80, 81, 90)
}

func TestNthUgly(t *testing.T) {
  nth := gomath.NewNthBigInt(gomath.Ugly(3, 5, 7))
  value := new(big.Int)
  assertBigIntEqual(t, 2401, nth.Nth(50, value))
  assertBigIntEqual(t, 33075, nth.Nth(100, value))
}

func TestSingleFactor(t *testing.T) {
  uglies := gomath.Ugly(3)
  checkInfBigIntStream(
      t,
      uglies,
      1, 3, 9, 27, 81, 243, 729, 2187, 6561, 19683, 59049)
}

func TestPanic(t *testing.T) {
  assertPanic(
      t,
      func() {
        gomath.Ugly(1, 2, 3)
      })
}

func BenchmarkUgly(b *testing.B) {
  uglies := gomath.Ugly(2, 3, 5)
  value := new(big.Int)
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    uglies.Next(value)
  }
}
