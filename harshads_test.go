package gomath_test

import (
  "context"
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

func TestNthHarshad(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  indexer := gomath.NewIntIndexer(gomath.Harshads(ctx, 0))
  assertEqual(t, int64(100), indexer.Nth(33))
  assertEqual(t, int64(372), indexer.Nth(100))
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
  for _ = range harshads {
  }
}
