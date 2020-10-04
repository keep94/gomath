package gomath_test

import (
	"testing"

	"github.com/keep94/gomath"
	asserts "github.com/stretchr/testify/assert"
)

func TestProducts(t *testing.T) {
	products := gomath.Products(6, 2)
	checkInt64Stream(
		t,
		products,
		1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18, 20, 24, 25, 30, 36)
}

func TestZeroProducts(t *testing.T) {
	products := gomath.Products(8, 0)
	checkInt64Stream(t, products, 1)
}

func TestOneProduct(t *testing.T) {
	products := gomath.Products(4, 1)
	checkInt64Stream(t, products, 1, 2, 3, 4)
}

func TestProductOfOne(t *testing.T) {
	products := gomath.Products(1, 4)
	checkInt64Stream(t, products, 1)
}

func TestProductsPanic(t *testing.T) {
	assert := asserts.New(t)
	assert.Panics(func() {
		gomath.Products(0, 4)
	})
	assert.Panics(func() {
		gomath.Products(5, -1)
	})
}
