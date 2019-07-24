package gomath_test

import (
  "context"
  "math"
  "testing"

  "github.com/keep94/gomath"
)

func TestPrimes(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  primes := gomath.Primes(ctx, 2)
  checkInfInt64Chan(
      t,
      primes,
      2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61,
      67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137)
}

func TestPrimesBigStart(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  primes := gomath.Primes(ctx, 991)
  checkInfInt64Chan(
      t,
      primes,
      991, 997, 1009, 1013, 1019, 1021, 1031, 1033, 1039)
  ch := gomath.NewIntChan(gomath.Primes(ctx, 1000))
  p1229 := ch.Nth(1229-168)
  p1230 := ch.Nth(1230-168)
  assertTrue(t, p1229 < 10000)
  assertTrue(t, p1230 > 10000)
}

func TestPrimesMax(t *testing.T) {
  primes := gomath.Primes(context.Background(), math.MaxInt64)
  for range primes {
    t.Error("Didn't expect any primes")
  }
}
  
/*
// Commented out for now as it takes 39 seconds to run this one test
func TestPrimesMax1000(t *testing.T) {
  start := int64(math.MaxInt64 - 1000)
  primes := gomath.Primes(context.Background(), start)
  primesFound := false
  for p := range primes {
    assertTrue(t, p >= start)
    primesFound = true
  }
  assertTrue(t, primesFound)
}
*/

func TestNthPrime(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  ch := gomath.NewIntChan(gomath.Primes(ctx, 0))
  assertEqual(t, int64(997), ch.Nth(168))
  p1229 := ch.Nth(1229)
  p1230 := ch.Nth(1230)
  assertTrue(t, p1229 < 10000)
  assertTrue(t, p1230 > 10000)
  p9592 := ch.Nth(9592)
  p9593 := ch.Nth(9593)
  assertTrue(t, p9592 < 100000)
  assertTrue(t, p9593 > 100000)
}

func TestDecadePrimes(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  decades := gomath.DecadePrimes(ctx, 0)
  checkInfInt64Chan(
      t,
      decades,
      1, 10, 19, 82, 148, 187, 208)
}

func TestDecadePrimesBigStart(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  decades := gomath.DecadePrimes(ctx, 10)
  checkInfInt64Chan(
      t,
      decades,
      10, 19, 82, 148, 187, 208)
}

func BenchmarkPrimes(b *testing.B) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  primes := gomath.Primes(ctx, 2)
  i := 0
  b.ResetTimer()
  for range primes {
    i++
    if i == b.N {
      break
    }
  }
}

func TestPrimesContext(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  primes := gomath.Primes(ctx, 2)
  cancel()
  for range primes {
  }
}

func TestDecadePrimesContext(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  decades := gomath.DecadePrimes(ctx, 100)
  cancel()
  for range decades {
  }
}
