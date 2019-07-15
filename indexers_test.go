package gomath_test

import (
  "math/big"
  "testing"

  "github.com/keep94/gomath"
)

func TestBigIntIndexer(t *testing.T) {
  indexer := gomath.NewBigIntIndexer(upTo45By3())
  result := new(big.Int)
  assertPanic(t, func() {
    indexer.Nth(0, result)
  })
  assertBigIntEqual(t, 6, indexer.Nth(2, result))
  assertTrue(t, result == indexer.Nth(2, result))
  assertBigIntEqual(t, 45, indexer.Nth(15, result))
  assertPanic(t, func() {
    indexer.Nth(16, result)
  })
  assertBigIntEqual(t, 30, indexer.Nth(10, result))
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
  assertPanic(t, func() {
    indexer.Nth(0)
  })
  assertEqual(t, int64(6),indexer.Nth(2))
  assertEqual(t, int64(45), indexer.Nth(15))
  assertPanic(t, func() {
    indexer.Nth(16)
  })
  assertEqual(t, int64(30), indexer.Nth(10))
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
