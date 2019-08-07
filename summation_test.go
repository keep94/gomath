package gomath_test

import (
  "context"
  "testing"

  "github.com/keep94/gomath"
)

func TestSummation(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  sumPrimes := gomath.Summation(ctx, gomath.Primes(ctx, 2))
  checkInfInt64Chan(
      t,
      sumPrimes,
      2, 5, 10, 17, 28, 41, 58, 77, 100, 129, 160, 197, 238, 281, 328, 381)
}

func TestSummationContext(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  sumPrimes := gomath.Summation(ctx, gomath.Primes(ctx, 2))
  cancel()
  for range sumPrimes {
  }
}
