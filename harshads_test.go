package gomath_test

import (
  "context"
  "math"
  "testing"

  "github.com/keep94/gomath"
)

func TestHarshads(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  harshads := gomath.Harshads(ctx, 90)
  checkInfInt64Chan(
      t,
      harshads,
      90, 100, 102, 108, 110, 111, 112, 114, 117, 120)
}

func TestHarshadsMax(t *testing.T) {
  start := int64(math.MaxInt64 - 1000)
  harshads := gomath.Harshads(context.Background(), start)
  found := false
  for h := range harshads {
    assertTrue(t, h >= start)
    found = true
  }
  assertTrue(t, found)
}

func TestNthHarshad(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  ch := gomath.NewIntChan(gomath.Harshads(ctx, 0))
  assertEqual(t, int64(100), ch.Nth(33))
  assertEqual(t, int64(372), ch.Nth(100))
}

func BenchmarkHarshads(b *testing.B) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  harshads := gomath.Harshads(ctx, 1)
  i := 0
  b.ResetTimer()
  for range harshads {
    i++
    if i == b.N {
      break
    }
  }
}
    
func TestHarshadsContext(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  harshads := gomath.Harshads(ctx, 1)
  cancel()
  for range harshads {
  }
}
