package gomath_test

import (
	"testing"

	"github.com/keep94/gomath"
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
