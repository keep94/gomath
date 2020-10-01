package gomath

import (
	"github.com/keep94/gocombinatorics"
	"sort"
)

// Products generates all the numbers that can be written as a product of
// count positive integers each ranging between 1 and n.
func Products(n, count int) IntStream {
	stream := gocombinatorics.CombinationsWithReplacement(n, count)
	values := make([]int, stream.TupleSize())
	var products []int64
	for stream.Next(values) {
		products = append(products, computeProduct(values))
	}
	sort.Slice(
		products, func(i, j int) bool { return products[i] < products[j] })
	return &sliceStream{values: products}
}

func computeProduct(values []int) int64 {
	result := int64(1)
	for _, value := range values {
		result *= int64(value) + 1
	}
	return result
}

type sliceStream struct {
	values []int64
	index  int
}

func (s *sliceStream) Next() (int64, bool) {
	if s.index == len(s.values) {
		return 0, false
	}
	result := s.values[s.index]
	s.index++
	for s.index < len(s.values) && result == s.values[s.index] {
		s.index++
	}
	return result, true
}
