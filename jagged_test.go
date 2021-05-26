package gomath_test

import (
	"testing"

	"github.com/keep94/gomath"
	"github.com/stretchr/testify/assert"
)

func TestNormal(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(0.0, gomath.Jagged(0.0))
	assert.Equal(0.0, gomath.Jagged(2.0))
	assert.Equal(0.625, gomath.Jagged(-1.375))
	assert.Equal(0.5, gomath.Jagged(0.5))
	assert.Equal(0.625, gomath.Jagged(0.375))
	assert.InEpsilon(0.666667, gomath.Jagged(1.0/3.0), 1e-6)
	assert.InEpsilon(0.666667, gomath.Jagged(2.0/5.0), 1e-6)
}
