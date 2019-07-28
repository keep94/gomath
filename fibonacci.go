package gomath

import (
  "context"
  "math/big"
)

// Fibonacci generates fibonacci numbers. first and second are the
// first and second terms in the sequence, normally 1 and 1. 
func Fibonacci(ctx context.Context, first, second int64) <-chan *big.Int {
  result := make(chan *big.Int)
  a := big.NewInt(first)
  b := big.NewInt(second)
  go func() {
    defer close(result)
    for {
      select {
        case <-ctx.Done():
          return
        case result <- new(big.Int).Set(a):
      }
      a.Add(a, b)
      a, b = b, a
    }
  }()
  return result
}
