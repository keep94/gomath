package gomath

import (
  "context"
)

// Primes generates the prime numbers in order that are greater than or
// equal to start.
func Primes(ctx context.Context, start int64) <-chan int64 {
  result := make(chan int64)
  go func() {
    defer close(result)
    if start <= 2 {
      select {
        case <-ctx.Done():
          return
        case result <- 2:
      }
      start = 3
    } else {
      start = start / 2 * 2 + 1  // start on odd
    }
    sieveSize := int64(1000)
    for {
      sieve := initSieve(int(sieveSize))
      factor := int64(3)
      for factor*factor < start + sieveSize*2 {
        var mult int64
        if factor*factor >= start {
          mult = factor * factor
        } else {
          // smallest odd multiple of factor that is >= start
          mult = (start + factor - 1) / (2*factor) * (2*factor) + factor
        }
        for mult < start + sieveSize*2 {
          sieve[int((mult - start) / 2)] = false
          mult += (2 * factor)
        }
        factor += 2
      }
      idx := int64(0)
      for idx < sieveSize {
        if sieve[int(idx)] {
          select {
            case <-ctx.Done():
              return
            case result <- start + 2*idx:
          }
        }
        idx++
      }
      start += sieveSize*2
      if sieveSize < 1000000 {
        sieveSize *= 2
      }
    }
  }()
  return result
}

func initSieve(sieveSize int) []bool {
  result := make([]bool, sieveSize)
  for i := range result {
    result[i] = true
  }
  return result
}

// DecadePrimes generates all x >= start in ascending order such that
// 10x + 1, 10x + 3, 10x + 7, and 10x + 9 are all prime.
func DecadePrimes(ctx context.Context, start int64) <-chan int64 {
  if start < 1 {
    start = 1
  }
  lastDecade := int64(0)
  primeCount := 0
  primes := Primes(ctx, 10*start + 1)
  result := make(chan int64)
  go func() {
    defer close(result)
    for p := range primes {
      if p / 10 == lastDecade {
        primeCount++
      } else {
        lastDecade = p / 10
        primeCount = 1
      }
      if primeCount == 4 {
        select {
          case <-ctx.Done():
            return
          case result <- lastDecade:
        }
      }
    }
  }()
  return result
}
