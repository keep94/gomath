package gomath_test

import (
  "math/big"
  "testing"

  "github.com/keep94/gomath"
)

func TestBigIntChan(t *testing.T) {
  ch := gomath.NewBigIntChan(upTo45By3())
  assertPanic(t, func() {
    ch.Nth(0)
  })
  assertBigIntEqual(t, 6, ch.Nth(2))
  assertPanic(t, func() {
    ch.Nth(2)
  })
  assertBigIntEqual(t, 45, ch.Nth(15))
  assertPanic(t, func() {
    ch.Nth(16)
  })
  assertPanic(t, func() {
    ch.Nth(10)
  })
}

func TestIntChan(t *testing.T) {
  ch := gomath.NewIntChan(upTo45By3Int())
  assertPanic(t, func() {
    ch.Nth(0)
  })
  assertEqual(t, int64(6), ch.Nth(2))
  assertPanic(t, func() {
    ch.Nth(2)
  })
  assertEqual(t, int64(45), ch.Nth(15))
  assertPanic(t, func() {
    ch.Nth(16)
  })
  assertPanic(t, func() {
    ch.Nth(10)
  })
}

func TestIntSafeChan(t *testing.T) {
  ch := gomath.NewIntChan(upTo45By3Int())
  assertPanic(t, func() {
    ch.SafeNth(0)
  })
  result, ok := ch.SafeNth(2)
  assertTrue(t, ok)
  assertEqual(t, int64(6), result)
  assertPanic(t, func() {
    ch.SafeNth(2)
  })
  result, ok = ch.SafeNth(16)
  assertTrue(t, !ok)
  assertEqual(t, int64(0), result)
  assertPanic(t, func() {
    ch.Nth(10)
  })
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
