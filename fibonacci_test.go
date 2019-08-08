package gomath_test

import (
  "context"
  "testing"

  "github.com/keep94/gomath"
)

func TestFibonacci(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  fib := gomath.Fibonacci(ctx, 1, 1)
  checkInfBigIntChan(
      t,
      fib,
      1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987)
}

func TestFibonacciDifferentTerms(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  fib := gomath.Fibonacci(ctx, 2, 10)
  checkInfBigIntChan(
      t,
      fib,
      2, 10, 12, 22, 34, 56, 90, 146, 236, 382, 618, 1000, 1618, 2618)
}

func BenchmarkFibonacci(b *testing.B) {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  fib := gomath.Fibonacci(ctx, 1, 1)
  i := 0
  b.ResetTimer()
  for range fib {
    i++
    if i == b.N {
      break
    }
  }
}
    
func TestFibonacciContext(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  fib := gomath.Fibonacci(ctx, 1, 1)
  cancel()
  for range fib {
  }
}
