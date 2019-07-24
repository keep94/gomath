package gomath

import (
  "context"
  "math"
)

// Harshads generates the harshad numbers in order that are greater than or
// equal to start.
func Harshads(ctx context.Context, start int64) <-chan int64 {
  if start < 1 {
    start = 1
  }
  sum := sumDigits(start)
  result := make(chan int64)
  go func() {
    defer close(result)
    for {
      if start % sum == 0 {
        select {
          case <-ctx.Done():
            return
          case result <- start:
        }
      }
      if start == math.MaxInt64 {
        return
      }
      start++
      sum++
      temp := start
      for temp % 10 == 0 {
        temp /= 10
        sum -= 9
      }
    }
  }()
  return result
}

func sumDigits(n int64) int64 {
  result := int64(0)
  for n > 0 {
    result += n % 10
    n /= 10
  }
  return result
}
