package gomath_test

import (
  "context"
  "testing"

  "github.com/keep94/gomath"
)

func TestUgly(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  uglies := gomath.Ugly(ctx, 2, 3, 5)
  checkInfBigIntChan(
      t,
      uglies,
      1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18, 20, 24, 25, 27,
      30, 32, 36, 40, 45, 48, 50, 54, 60, 64, 72, 75, 80, 81, 90)
}

func TestNthUgly(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  ch := gomath.NewBigIntChan(gomath.Ugly(ctx, 3, 5, 7))
  assertBigIntEqual(t, 2401, ch.Nth(50))
  assertBigIntEqual(t, 33075, ch.Nth(100))
}

func TestSingleFactor(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  uglies := gomath.Ugly(ctx, 3)
  checkInfBigIntChan(
      t,
      uglies,
      1, 3, 9, 27, 81, 243, 729, 2187, 6561, 19683, 59049)
}

func TestPanic(t *testing.T) {
  assertPanic(
      t,
      func() {
        gomath.Ugly(context.Background(), 1, 2, 3)
      })
}

func BenchmarkUgly(b *testing.B) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  uglies := gomath.Ugly(ctx, 2, 3, 5)
  i := 0
  b.ResetTimer()
  for range uglies {
    i++
    if i == b.N {
      break
    }
  }
}
    
func TestContext(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  uglies := gomath.Ugly(ctx, 3, 5, 7)
  cancel()
  for _ = range uglies {
  }
}
