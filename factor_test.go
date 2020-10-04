package gomath_test

import (
	"testing"

	"github.com/keep94/gomath"
	asserts "github.com/stretchr/testify/assert"
)

func TestFactor(t *testing.T) {
	assertPP(t, gomath.Factor(1))
	assertPP(t, gomath.Factor(2), 2, 1)
	assertPP(t, gomath.Factor(3), 3, 1)
	assertPP(t, gomath.Factor(4), 2, 2)
	assertPP(t, gomath.Factor(5), 5, 1)
	assertPP(t, gomath.Factor(6), 2, 1, 3, 1)
	assertPP(t, gomath.Factor(7), 7, 1)
	assertPP(t, gomath.Factor(8), 2, 3)
	assertPP(t, gomath.Factor(9), 3, 2)
	assertPP(t, gomath.Factor(10), 2, 1, 5, 1)
	assertPP(t, gomath.Factor(11), 11, 1)
	assertPP(t, gomath.Factor(12), 2, 2, 3, 1)
	assertPP(t, gomath.Factor(13), 13, 1)
	assertPP(t, gomath.Factor(14), 2, 1, 7, 1)
	assertPP(t, gomath.Factor(15), 3, 1, 5, 1)
	assertPP(t, gomath.Factor(16), 2, 4)
	assertPP(t, gomath.Factor(17), 17, 1)
	assertPP(t, gomath.Factor(18), 2, 1, 3, 2)
	assertPP(t, gomath.Factor(19), 19, 1)
	assertPP(t, gomath.Factor(20), 2, 2, 5, 1)
	assertPP(t, gomath.Factor(77), 7, 1, 11, 1)
	assertPP(t, gomath.Factor(86), 2, 1, 43, 1)
	assertPP(t, gomath.Factor(100), 2, 2, 5, 2)
	assertPP(t, gomath.Factor(2019), 3, 1, 673, 1)
	assertPP(t, gomath.Factor(10013), 17, 1, 19, 1, 31, 1)
	assertPP(t, gomath.Factor(10080), 2, 5, 3, 2, 5, 1, 7, 1)
	assertPanic(t, func() { gomath.Factor(0) })
}

func TestFactors(t *testing.T) {
	assert := asserts.New(t)
	assert.Equal([]int64{1}, gomath.Factors(1))
	assert.Equal([]int64{1, 2}, gomath.Factors(2))
	assert.Equal([]int64{1, 3}, gomath.Factors(3))
	assert.Equal([]int64{1, 2, 4}, gomath.Factors(4))
	assert.Equal([]int64{1, 2, 3, 6}, gomath.Factors(6))
	assert.Equal([]int64{1, 2, 3, 5, 6, 10, 15, 30}, gomath.Factors(30))
	assert.Equal(
		[]int64{1, 2, 3, 4, 6, 8, 9, 12, 18, 24, 36, 72},
		gomath.Factors(72))
	assert.Equal(
		[]int64{
			1,
			2,
			3,
			4,
			5,
			6,
			10,
			12,
			15,
			20,
			25,
			30,
			50,
			60,
			75,
			100,
			150,
			300},
		gomath.Factors(300))
	assert.Panics(func() {
		gomath.Factors(0)
	})
}
