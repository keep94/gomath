package gomath_test

import (
	"math/big"
	"testing"

	"github.com/keep94/gomath"
)

func TestFibonacci(t *testing.T) {
	fib := gomath.Fibonacci(1, 1)
	checkInfBigIntStream(
		t,
		fib,
		1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987)
}

func TestFibonacciDifferentTerms(t *testing.T) {
	fib := gomath.Fibonacci(2, 10)
	checkInfBigIntStream(
		t,
		fib,
		2, 10, 12, 22, 34, 56, 90, 146, 236, 382, 618, 1000, 1618, 2618)
}

func BenchmarkFibonacci(b *testing.B) {
	fib := gomath.Fibonacci(1, 1)
	value := new(big.Int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fib.Next(value)
	}
}
