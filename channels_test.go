package gomath_test

import (
  "math"
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

func TestBigIntCounter(t *testing.T) {
  counter := gomath.NewBigIntCounter(func(x int64) *big.Int {
    return big.NewInt(3*x)
  })
  value := new(big.Int)
  assertEqual(t, int64(0), counter.CountLE(value.SetInt64(math.MinInt64)))
  assertEqual(t, int64(0), counter.CountLE(value.SetInt64(2)))
  assertEqual(t, int64(1), counter.CountLE(value.SetInt64(3)))
  assertEqual(t, int64(1), counter.CountLE(value.SetInt64(3)))
  assertPanic(t, func() {
    counter.CountLE(value.SetInt64(2))
  })
  assertEqual(t, int64(1), counter.CountLE(value.SetInt64(5)))
  assertEqual(t, int64(2), counter.CountLE(value.SetInt64(6)))
  assertEqual(t, int64(5), counter.CountLE(value.SetInt64(17)))
  assertEqual(t, int64(5), counter.CountLE(value.SetInt64(17)))
  assertEqual(t, int64(6), counter.CountLE(value.SetInt64(18)))
  assertEqual(t, int64(14), counter.CountLE(value.SetInt64(44)))
  assertEqual(t, int64(15), counter.CountLE(value.SetInt64(45)))
  assertEqual(t, int64(33), counter.CountLE(value.SetInt64(100)))
}

func TestIntCounter(t *testing.T) {
  counter := gomath.NewIntCounter(func(x int64) int64 {
    if x <= 15 {
      return 3*x
    }
    return math.MaxInt64
  })
  assertEqual(t, int64(0), counter.CountLE(math.MinInt64))
  assertEqual(t, int64(0), counter.CountLE(2))
  assertEqual(t, int64(1), counter.CountLE(3))
  assertEqual(t, int64(1), counter.CountLE(3))
  assertPanic(t, func() {
    counter.CountLE(2)
  })
  assertEqual(t, int64(1), counter.CountLE(5))
  assertEqual(t, int64(2), counter.CountLE(6))
  assertEqual(t, int64(5), counter.CountLE(17))
  assertEqual(t, int64(5), counter.CountLE(17))
  assertEqual(t, int64(6), counter.CountLE(18))
  assertEqual(t, int64(14), counter.CountLE(44))
  assertEqual(t, int64(15), counter.CountLE(45))
  assertEqual(t, int64(15), counter.CountLE(100))
  assertEqual(t, int64(16), counter.CountLE(math.MaxInt64))
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
