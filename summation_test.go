package gomath_test

import (
  "testing"

  "github.com/keep94/gomath"
)

func TestIntSummation(t *testing.T) {
  sumPrimes := gomath.IntSummation(gomath.Primes(2))
  checkInfInt64Stream(
      t,
      sumPrimes,
      2, 5, 10, 17, 28, 41, 58, 77, 100, 129, 160, 197, 238, 281, 328, 381)
}

func TestBigIntSummation(t *testing.T) {
  sumPartitions := gomath.BigIntSummation(gomath.Partitions())
  checkInfBigIntStream(
      t,
      sumPartitions,
      1, 3, 6, 11, 18, 29, 44, 66, 96, 138, 194, 271, 372)
}
