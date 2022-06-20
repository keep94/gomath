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
	products := make(map[int64]struct{})
	for stream.Next(values) {
		products[computeProduct(values)] = struct{}{}
	}
	result := make([]int64, 0, len(products))
	for p := range products {
		result = append(result, p)
	}
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	return result
}

// Deprecated: Use ProductsSlice.
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
