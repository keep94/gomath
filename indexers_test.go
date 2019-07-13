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

func TestBigIntIndexerNil(t *testing.T) {
  indexer := gomath.NewBigIntIndexer(nilBigInts())
  result := new(big.Int)
  assertPanic(t, func() {
    indexer.Nth(5, result)
  })
}

func TestIntIndexer(t *testing.T) {
  indexer := gomath.NewIntIndexer(upTo45By3Int())
  _, ok := indexer.Nth(0)
  assertFalse(t, ok)
  n, ok := indexer.Nth(2)
  assertEqual(t, int64(6), n)
  assertTrue(t, ok)
  n, ok = indexer.Nth(15)
  assertEqual(t, int64(45), n)
  assertTrue(t, ok)
  _, ok = indexer.Nth(18)
  assertFalse(t, ok)
  _, ok = indexer.Nth(17)
  assertFalse(t, ok)
  _, ok = indexer.Nth(16)
  assertFalse(t, ok)
  n, ok = indexer.Nth(10)
  assertEqual(t, int64(30), n)
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

func nilBigInts() <-chan *big.Int {
  result := make(chan *big.Int)
  go func() {
    defer close(result)
    result <- big.NewInt(3)
    result <- big.NewInt(3)
    result <- nil
    result <- big.NewInt(3)
    result <- big.NewInt(3)
  }()
  return result
}

func upTo45By3Int() <-chan int64 {
  result := make(chan int64)
  go func() {
    defer close(result)
    for i := int64(1); i <= 15; i++ {
      result <- 3*i
    }
  }()
  return result
}
