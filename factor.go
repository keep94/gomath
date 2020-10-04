package gomath

import (
	"sort"
)

// PrimePower represents a single term of a prime power decomposition
type PrimePower struct {
	Prime int64
	Power int
}

// Factor returns the prime power decomposition of n. Factor panics if n < 1.
func Factor(n int64) []PrimePower {
	if n < 1 {
		panic("n can't be less than 1")
	}
	fact := int64(2)
	var result []PrimePower
	for fact <= n/fact {
		exp := 0
		for n%fact == 0 {
			n /= fact
			exp++
		}
		if exp > 0 {
			result = append(result, PrimePower{Prime: fact, Power: exp})
		}
		fact++
	}
	if n > 1 {
		result = append(result, PrimePower{Prime: n, Power: 1})
	}
	return result
}

// Factors returns all the positive integers that divide n in ascending order.
// Factors panics if n < 1.
func Factors(n int64) []int64 {
	return ppdFactors(Factor(n))
}

func ppdFactorCount(ppd []PrimePower) int {
	result := 1
	for _, pp := range ppd {
		result *= (pp.Power + 1)
	}
	return result
}

func ppdFactors(ppd []PrimePower) []int64 {
	result := make([]int64, ppdFactorCount(ppd))
	result[0] = 1
	length := 1
	for _, pp := range ppd {
		for i := 0; i < length*pp.Power; i++ {
			result[length+i] = result[i] * pp.Prime
		}
		length *= (pp.Power + 1)
	}
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	return result
}
