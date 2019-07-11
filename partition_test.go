package gomath_test

import (
  "math/big"
  "testing"
  "github.com/keep94/gomath"
)

func TestPartition(t *testing.T) {
  p := gomath.NewPartition()
  assertBigIntEqual(t, 1, p.Eval(0, new(big.Int)))
  assertBigIntEqual(t, 7, p.Eval(5, new(big.Int)))
  assertBigIntEqual(t, 42, p.Eval(10, new(big.Int)))
  assertBigIntEqual(t, 627, p.Eval(20, new(big.Int)))
  assertBigIntEqual(t, 5604, p.Eval(30, new(big.Int)))
  assertBigIntEqual(t, 5604, p.Chart(30))
  assertBigIntEqual(t, 37338, p.Eval(40, new(big.Int)))
  assertBigIntEqual(t, 101, p.Eval(13, new(big.Int)))
  assertBigIntEqual(t, 190569292, p.Eval(100, new(big.Int)))
  assertPanic(t, func() {
    p.Eval(-1, new(big.Int))
  })
  result := new(big.Int)
  assertTrue(t, result == p.Eval(7, result))
  assertBigIntEqual(t, 15, result)
}
