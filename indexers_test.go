package gomath_test

import (
  "math/big"
  "testing"

  "github.com/keep94/gomath"
)

func TestBigIntIndexer(t *testing.T) {
  indexer := gomath.NewBigIntIndexer(upTo45By3())
  result := new(big.Int)
  n, ok := indexer.Nth(0, result)
  assertBigIntNil(t, n)
  assertFalse(t, ok)
  n, ok = indexer.Nth(2, result)
  assertBigIntEqual(t, 6, result)
  assertTrue(t, n == result)
  assertTrue(t, ok)
  n, ok = indexer.Nth(15, result)
  assertBigIntEqual(t, 45, result)
  assertTrue(t, n == result)
  assertTrue(t, ok)
  n, ok = indexer.Nth(18, result)
  assertBigIntNil(t, n)
  assertFalse(t, ok)
  n, ok = indexer.Nth(17, result)
  assertBigIntNil(t, n)
  assertFalse(t, ok)
  n, ok = indexer.Nth(16, result)
  assertBigIntNil(t, n)
  assertFalse(t, ok)
  n, ok = indexer.Nth(10, result)
  assertBigIntEqual(t, 30, result)
  assertTrue(t, n == result)
  assertTrue(t, ok)
}

func upTo45By3() <-chan *big.Int {
  result := make(chan *big.Int)
  go func() {
    defer close(result)
    for i := 1; i <= 15; i++ {
      result <- big.NewInt(3*int64(i))
    }
  }()
  return result
}
