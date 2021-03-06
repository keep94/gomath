package gomath

import (
	"sort"

	"github.com/keep94/gocombinatorics"
)

// ProductsSlice returns all the numbers in ascending order that can be written
// as a product of count positive integers each ranging between 1 and n.
// ProductsSlice panics if n < 1 or if count is negative.
func ProductsSlice(n, count int) []int64 {
	if n < 1 || count < 0 {
		panic("n must be greater than 0 and count must be non negative.")
	}
	stream := gocombinatorics.CombinationsWithReplacement(n, count)
	values := make([]int, stream.TupleSize())
	var products []int64
	for stream.Next(values) {
		products = append(products, computeProduct(values))
	}
	sort.Slice(
		products, func(i, j int) bool { return products[i] < products[j] })
	idx := 1
	for i := 1; i < len(products); i++ {
		if products[i] == products[i-1] {
			continue
		}
		products[idx] = products[i]
		idx++
	}
	result := make([]int64, idx)
	copy(result, products)
	return result
}

// Products is deprecated in favor of ProductsSlice.
func Products(n, count int) IntStream {
	return &sliceIntStream{values: ProductsSlice(n, count)}
}

func computeProduct(values []int) int64 {
	result := int64(1)
	for _, value := range values {
		result *= int64(value) + 1
	}
	return result
}

type sliceIntStream struct {
	values []int64
	idx    int
}

func (s *sliceIntStream) Next() (int64, bool) {
	if s.idx == len(s.values) {
		return 0, false
	}
	result := s.values[s.idx]
	s.idx++
	return result, true
}
