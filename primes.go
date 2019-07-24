package gomath

import (
  "context"
  "math"
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
      if start == math.MaxInt64 {
        return
      }
      if sieveSize*2 > math.MaxInt64 - start {
        sieveSize = (math.MaxInt64 - start) / 2
      }
      sieve := initSieve(int(sieveSize))
      factor := int64(3)
      for factor <= (start + sieveSize*2 - 2) / factor {
        multStart := divideRoundUpOdd(start, factor)
        if multStart < factor {
          multStart = factor
        }
        multEnd := divideRoundUpOdd(start + sieveSize*2, factor)
        for mult := multStart; mult < multEnd; mult += 2 {
          sieve[int((mult*factor - start) / 2)] = false
        }
        factor += 2
      }
      for idx := int64(0); idx < sieveSize; idx++ {
        if sieve[int(idx)] {
          select {
            case <-ctx.Done():
              return
            case result <- start + 2*idx:
          }
        }
      }
      start += sieveSize*2
      if sieveSize < 1000000 {
        sieveSize *= 2
      }
    }
  }()
  return result
}

func divideRoundUpOdd(n, d int64) int64 {
  roundUp := (n - 1) / d + 1
  return roundUp / 2 * 2 + 1
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
