package gomath_test

import (
  "math/big"
  "testing"

  "github.com/keep94/gomath"
)

func TestBigIntIndexer(t *testing.T) {
  indexer := gomath.NewBigIntIndexer(upTo45By3())
  n, ok := indexer.Nth(0)
  assertBigIntNil(t, n)
  assertFalse(t, ok)
  n, ok = indexer.Nth(2)
  assertBigIntEqual(t, 6, n)
  assertTrue(t, ok)
  n.Set(big.NewInt(50))
  n, _ = indexer.Nth(2)
  assertBigIntEqual(t, 6, n)
  n, ok = indexer.Nth(15)
  assertBigIntEqual(t, 45, n)
  assertTrue(t, ok)
  n, ok = indexer.Nth(18)
  assertBigIntNil(t, n)
  assertFalse(t, ok)
  n, ok = indexer.Nth(17)
  assertBigIntNil(t, n)
  assertFalse(t, ok)
  n, ok = indexer.Nth(16)
  assertBigIntNil(t, n)
  assertFalse(t, ok)
  n, ok = indexer.Nth(10)
  assertBigIntEqual(t, 30, n)
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
