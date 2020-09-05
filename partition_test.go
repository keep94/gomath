package gomath_test

import (
	"math/big"
	"testing"

	"github.com/keep94/gomath"
)

func TestPartition(t *testing.T) {
	p := gomath.NewPartition()
	value := new(big.Int)
	assertBigIntEqual(t, 1, p.Eval(0, value))
	assertBigIntEqual(t, 7, p.Eval(5, value))
	assertBigIntEqual(t, 42, p.Eval(10, value))
	assertBigIntEqual(t, 627, p.Eval(20, value))
	assertBigIntEqual(t, 5604, p.Eval(30, value))
	assertBigIntEqual(t, 5604, p.Chart(30, value))
	assertBigIntEqual(t, 37338, p.Eval(40, value))
	assertBigIntEqual(t, 101, p.Eval(13, value))
	assertBigIntEqual(t, 190569292, p.Eval(100, value))
	assertPanic(t, func() {
		p.Eval(-1, new(big.Int))
	})
	result := new(big.Int)
	assertTrue(t, result == p.Eval(7, result))
	assertBigIntEqual(t, 15, result)
}

func TestPartitions(t *testing.T) {
	stream := gomath.Partitions()
	checkInfBigIntStream(t, stream, 1, 2, 3, 5, 7, 11, 15, 22, 30, 42, 56, 77, 101)
}
