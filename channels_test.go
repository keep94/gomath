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

func TestInvBigIntChanEmpty(t *testing.T) {
  ch := make(chan *big.Int)
  close(ch)
  empty := gomath.NewInvBigIntChan(ch)
  value := new(big.Int)
  assertEqual(t, int64(0), empty.InvNth(value.SetInt64(math.MinInt64)))
  assertEqual(t, int64(0), empty.InvNth(value.SetInt64(math.MaxInt64)))
}

func TestInvIntChanEmpty(t *testing.T) {
  ch := make(chan int64)
  close(ch)
  empty := gomath.NewInvIntChan(ch)
  assertEqual(t, int64(0), empty.InvNth(math.MinInt64))
  assertEqual(t, int64(0), empty.InvNth(math.MaxInt64))
}

func TestInvBigIntChan(t *testing.T) {
  ch := gomath.NewInvBigIntChan(upTo45By3())
  value := new(big.Int)
  assertEqual(t, int64(0), ch.InvNth(value.SetInt64(math.MinInt64)))
  assertEqual(t, int64(0), ch.InvNth(value.SetInt64(2)))
  assertEqual(t, int64(1), ch.InvNth(value.SetInt64(3)))
  assertEqual(t, int64(1), ch.InvNth(value.SetInt64(3)))
  assertPanic(t, func() {
    ch.InvNth(value.SetInt64(2))
  })
  assertEqual(t, int64(1), ch.InvNth(value.SetInt64(5)))
  assertEqual(t, int64(2), ch.InvNth(value.SetInt64(6)))
  assertEqual(t, int64(5), ch.InvNth(value.SetInt64(17)))
  assertEqual(t, int64(5), ch.InvNth(value.SetInt64(17)))
  assertEqual(t, int64(6), ch.InvNth(value.SetInt64(18)))
  assertEqual(t, int64(14), ch.InvNth(value.SetInt64(44)))
  assertEqual(t, int64(15), ch.InvNth(value.SetInt64(45)))
  assertEqual(t, int64(15), ch.InvNth(value.SetInt64(100)))
}

func TestInvIntChan(t *testing.T) {
  ch := gomath.NewInvIntChan(upTo45By3Int())
  assertEqual(t, int64(0), ch.InvNth(math.MinInt64))
  assertEqual(t, int64(0), ch.InvNth(2))
  assertEqual(t, int64(1), ch.InvNth(3))
  assertEqual(t, int64(1), ch.InvNth(3))
  assertPanic(t, func() {
    ch.InvNth(2)
  })
  assertEqual(t, int64(1), ch.InvNth(5))
  assertEqual(t, int64(2), ch.InvNth(6))
  assertEqual(t, int64(5), ch.InvNth(17))
  assertEqual(t, int64(5), ch.InvNth(17))
  assertEqual(t, int64(6), ch.InvNth(18))
  assertEqual(t, int64(14), ch.InvNth(44))
  assertEqual(t, int64(15), ch.InvNth(45))
  assertEqual(t, int64(15), ch.InvNth(100))
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
