package gomath_test

import (
	"math"
	"testing"

	"github.com/keep94/gomath"
)

func TestPrimes(t *testing.T) {
	primes := gomath.Primes(2)
	checkInfInt64Stream(
		t,
		primes,
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61,
		67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137)
}

func TestPrimesBigStart(t *testing.T) {
	primes := gomath.Primes(991)
	checkInfInt64Stream(
		t,
		primes,
		991, 997, 1009, 1013, 1019, 1021, 1031, 1033, 1039)
	nth := gomath.NewNthInt(gomath.Primes(1000))
	p1229 := nth.Nth(1229 - 168)
	p1230 := nth.Nth(1230 - 168)
	assertTrue(t, p1229 < 10000)
	assertTrue(t, p1230 > 10000)
}

func TestPrimesMax(t *testing.T) {
	primes := gomath.Primes(math.MaxInt64)
	_, ok := primes.Next()
	if ok {
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
	nth := gomath.NewNthInt(gomath.Primes(0))
	assertEqual(t, int64(997), nth.Nth(168))
	p1229 := nth.Nth(1229)
	p1230 := nth.Nth(1230)
	assertTrue(t, p1229 < 10000)
	assertTrue(t, p1230 > 10000)
	p9592 := nth.Nth(9592)
	p9593 := nth.Nth(9593)
	assertTrue(t, p9592 < 100000)
	assertTrue(t, p9593 > 100000)
	p78498 := nth.Nth(78498)
	p78499 := nth.Nth(78499)
	assertTrue(t, p78498 < 1000000)
	assertTrue(t, p78499 > 1000000)
	p664579 := nth.Nth(664579)
	p664580 := nth.Nth(664580)
	assertTrue(t, p664579 < 10000000)
	assertTrue(t, p664580 > 10000000)
}

func TestDecadePrimes(t *testing.T) {
	decades := gomath.DecadePrimes(0)
	checkInfInt64Stream(
		t,
		decades,
		1, 10, 19, 82, 148, 187, 208)
}

func TestDecadePrimesBigStart(t *testing.T) {
	decades := gomath.DecadePrimes(10)
	checkInfInt64Stream(
		t,
		decades,
		10, 19, 82, 148, 187, 208)
}

func BenchmarkPrimes(b *testing.B) {
	primes := gomath.Primes(2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		primes.Next()
	}
}
