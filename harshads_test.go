package gomath_test

import (
  "math"
  "testing"

  "github.com/keep94/gomath"
)

func TestHarshads(t *testing.T) {
  harshads := gomath.Harshads(90)
  checkInfInt64Stream(
      t,
      harshads,
      90, 100, 102, 108, 110, 111, 112, 114, 117, 120)
}

func TestHarshadsMax(t *testing.T) {
  start := int64(math.MaxInt64 - 1000)
  harshads := gomath.Harshads(start)
  found := false
  harshad, ok := harshads.Next()
  for ; ok; harshad, ok = harshads.Next() {
    assertTrue(t, harshad >= start)
    found = true
  }
  assertTrue(t, found)
}

func TestNthHarshad(t *testing.T) {
  nth := gomath.NewNthInt(gomath.Harshads(0))
  assertEqual(t, int64(100), nth.Nth(33))
  assertEqual(t, int64(372), nth.Nth(100))
}

func BenchmarkHarshads(b *testing.B) {
  harshads := gomath.Harshads(1)
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    harshads.Next()
  }
}
