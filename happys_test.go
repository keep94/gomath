package gomath_test

import (
  "context"
  "math"
  "testing"

  "github.com/keep94/gomath"
)

func TestHappys(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  happys := gomath.Happys(ctx, 100)
  checkInfInt64Chan(
      t,
      happys,
      100, 103, 109, 129, 130, 133, 139, 167, 176, 188)
}

func TestHappysMax(t *testing.T) {
  start := int64(math.MaxInt64 - 1000)
  happys := gomath.Happys(context.Background(), start)
  found := false
  for h := range happys {
    assertTrue(t, h >= start)
    found = true
  }
  assertTrue(t, found)
}

func TestNthHappy(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  ch := gomath.NewIntChan(gomath.Happys(ctx, 0))
  assertEqual(t, int64(100), ch.Nth(20))
  assertEqual(t, int64(694), ch.Nth(100))
  assertEqual(t, int64(6899), ch.Nth(1000))
  assertEqual(t, int64(67169), ch.Nth(10000))
}

func BenchmarkHappys(b *testing.B) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  happys := gomath.Happys(ctx, 1)
  i := 0
  b.ResetTimer()
  for range happys {
    i++
    if i == b.N {
      break
    }
  }
}

func TestHappysContext(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  happys := gomath.Happys(ctx, 1)
  cancel()
  for range happys {
  }
}
