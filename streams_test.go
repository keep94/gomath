package gomath_test

import (
	"math"
	"math/big"
	"testing"

	"github.com/keep94/gomath"
)

func TestNthBigInt(t *testing.T) {
	nth := gomath.NewNthBigInt(newUpBy3Stream())
	value := new(big.Int)
	assertPanic(t, func() {
		nth.Nth(0, value)
	})
	assertBigIntEqual(t, 6, nth.Nth(2, value))
	assertPanic(t, func() {
		nth.Nth(2, value)
	})
	assertBigIntEqual(t, 45, nth.Nth(15, value))
	assertPanic(t, func() {
		nth.Nth(10, value)
	})
}

func TestNthInt(t *testing.T) {
	nth := gomath.NewNthInt(newUpTo45By3Stream())
	assertPanic(t, func() {
		nth.Nth(0)
	})
	assertEqual(t, int64(6), nth.Nth(2))
	assertPanic(t, func() {
		nth.Nth(2)
	})
	assertEqual(t, int64(45), nth.Nth(15))
	assertPanic(t, func() {
		nth.Nth(16)
	})
	assertPanic(t, func() {
		nth.Nth(10)
	})
}

func TestNthIntSafe(t *testing.T) {
	nth := gomath.NewNthInt(newUpTo45By3Stream())
	assertPanic(t, func() {
		nth.SafeNth(0)
	})
	result, ok := nth.SafeNth(2)
	assertTrue(t, ok)
	assertEqual(t, int64(6), result)
	assertPanic(t, func() {
		nth.SafeNth(2)
	})
	result, ok = nth.SafeNth(16)
	assertTrue(t, !ok)
	assertEqual(t, int64(0), result)
	assertPanic(t, func() {
		nth.Nth(10)
	})
}

func TestBigIntCounter(t *testing.T) {
	counter := gomath.NewBigIntCounter(newUpBy3Stream())
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
	counter := gomath.NewIntCounter(newUpTo45By3Stream())
	assertEqual(t, int64(0), counter.CountLE(math.MinInt64))
	assertEqual(t, int64(0), counter.CountLE(2))
	assertEqual(t, int64(1), counter.CountLE(3))
	assertEqual(t, int64(1), counter.CountLE(3))
	assertPanic(t, func() {
		counter.CountLE(2)
	})
	assertEqual(t, int64(1), counter.CountLE(4))
	assertEqual(t, int64(1), counter.CountLE(5))
	assertEqual(t, int64(2), counter.CountLE(6))
	assertEqual(t, int64(5), counter.CountLE(17))
	assertEqual(t, int64(5), counter.CountLE(17))
	assertEqual(t, int64(6), counter.CountLE(18))
	assertEqual(t, int64(14), counter.CountLE(44))
	assertEqual(t, int64(15), counter.CountLE(45))
	assertEqual(t, int64(15), counter.CountLE(100))
	assertEqual(t, int64(15), counter.CountLE(math.MaxInt64))
}

func TestIntCounterEmptyStream(t *testing.T) {
	counter := gomath.NewIntCounter(newEmptyStream())
	assertEqual(t, int64(0), counter.CountLE(math.MinInt64))
	assertEqual(t, int64(0), counter.CountLE(math.MaxInt64))
}

type linearBigIntStream struct {
	currentValue *big.Int
	incrValue    *big.Int
}

func newUpBy3Stream() gomath.BigIntStream {
	return &linearBigIntStream{currentValue: big.NewInt(3), incrValue: big.NewInt(3)}
}

func (s *linearBigIntStream) Next(value *big.Int) *big.Int {
	value.Set(s.currentValue)
	s.currentValue.Add(s.currentValue, s.incrValue)
	return value
}

type linearIntStream struct {
	currentValue int64
	incrValue    int64
	maxValue     int64
}

func newUpTo45By3Stream() gomath.IntStream {
	return &linearIntStream{currentValue: 3, incrValue: 3, maxValue: 45}
}

func newEmptyStream() gomath.IntStream {
	return &linearIntStream{currentValue: 3, incrValue: 3, maxValue: 2}
}

func (s *linearIntStream) Next() (result int64, ok bool) {
	if s.currentValue > s.maxValue {
		return
	}
	result = s.currentValue
	ok = true
	s.currentValue += s.incrValue
	return
}
