package gomath_test

import (
  "context"
  "testing"

  "github.com/keep94/gomath"
)

func TestIntSummation(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  sumPrimes := gomath.IntSummation(ctx, gomath.Primes(ctx, 2))
  checkInfInt64Chan(
      t,
      sumPrimes,
      2, 5, 10, 17, 28, 41, 58, 77, 100, 129, 160, 197, 238, 281, 328, 381)
}

func TestIntSummationContext(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  sumPrimes := gomath.IntSummation(ctx, gomath.Primes(ctx, 2))
  cancel()
  for range sumPrimes {
  }
}

func TestBigIntSummation(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  sumPartitions := gomath.BigIntSummation(ctx, gomath.Partitions(ctx))
  checkInfBigIntChan(
      t,
      sumPartitions,
      1, 3, 6, 11, 18, 29, 44, 66, 96, 138, 194, 271, 372)
}

func TestBigIntSummationContext(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  sumPartitions := gomath.BigIntSummation(ctx, gomath.Partitions(ctx))
  cancel()
  for range sumPartitions {
  }
}
