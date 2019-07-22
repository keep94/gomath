package gomath_test

import (
  "math/big"
  "testing"

  "github.com/keep94/gomath"
)

func TestCake(t *testing.T) {
  c := gomath.NewCake()
  assertBigIntEqual(t, 1, c.Eval(0, 0, new(big.Int)))
  assertBigIntEqual(t, 1, c.Eval(0, 1, new(big.Int)))
  assertBigIntEqual(t, 2, c.Eval(1, 1, new(big.Int)))
  assertBigIntEqual(t, 16, c.Eval(4, 4, new(big.Int)))
  assertBigIntEqual(t, 31, c.Eval(4, 5, new(big.Int)))
  assertBigIntEqual(t, 57, c.Eval(4, 6, new(big.Int)))
  assertBigIntEqual(t, 4, c.Eval(4, 2, new(big.Int)))
  assertBigIntEqual(t, 26, c.Eval(3, 5, new(big.Int)))
  assertPanic(t, func() {
    c.Eval(-1, -1, new(big.Int))
  })
  result := new(big.Int)
  assertTrue(t, result == c.Eval(2, 3, result))
  assertBigIntEqual(t, 7, result)
}
