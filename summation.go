package gomath

import (
  "context"
)

// Summation generates the summation of the values in ch. For example if
// ch generates the primes, Summation would generate 2, 5, 10, 17, ...
func Summation(ctx context.Context, ch <-chan int64) <-chan int64 {
  result := make(chan int64)
  go func() {
    defer close(result)
    var sum int64
    for i := range ch {
      sum += i
      select {
        case <-ctx.Done():
          return
        case result <- sum:
      }
    }
  }()
  return result
}
