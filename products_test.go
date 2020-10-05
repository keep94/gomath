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

func TestProductsSlice(t *testing.T) {
	assert := asserts.New(t)
	assert.Equal(
		[]int64{1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18, 20, 24, 25, 30, 36},
		gomath.ProductsSlice(6, 2))
}

func TestZeroProducts(t *testing.T) {
	assert := asserts.New(t)
	assert.Equal([]int64{1}, gomath.ProductsSlice(8, 0))
}

func TestOneProduct(t *testing.T) {
	assert := asserts.New(t)
	assert.Equal([]int64{1, 2, 3, 4}, gomath.ProductsSlice(4, 1))
}

func TestProductOfOne(t *testing.T) {
	assert := asserts.New(t)
	assert.Equal([]int64{1}, gomath.ProductsSlice(1, 4))
}

func TestProductsSlicePanic(t *testing.T) {
	assert := asserts.New(t)
	assert.Panics(func() {
		gomath.ProductsSlice(0, 4)
	})
	assert.Panics(func() {
		gomath.ProductsSlice(5, -1)
	})
}
