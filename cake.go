package gomath

import (
	"math/big"
)

// Cake is used to compute the number of pieces you can get cutting an n
// dimensional cake k times. Cake instances are not safe to use with multiple
// goroutines
type Cake struct {
	values [][]*big.Int
	n      int
	k      int
}

// NewCake returns a new Cake instance
func NewCake() *Cake {
	n0 := []*big.Int{big.NewInt(1)}
	values := [][]*big.Int{n0}
	return &Cake{values: values}
}

// Eval stores the maximum  number of pieces you can get cutting an n
// dimensional cake k times in result and returns result. Eval panics if
// n < 0 or k < 0.
func (c *Cake) Eval(n, k int, result *big.Int) *big.Int {
	if n < 0 {
		panic("n must be greater than or equal to 0")
	}
	if k < 0 {
		panic("k must be greater than or equal to 0")
	}
	if k > c.k {
		for ni := 0; ni < c.n+1; ni++ {
			for ki := c.k + 1; ki < k+1; ki++ {
				if ni == 0 {
					// append 1
					c.values[0] = append(c.values[0], c.values[0][0])
				} else {
					c.values[ni] = append(
						c.values[ni],
						new(big.Int).Add(c.values[ni][ki-1], c.values[ni-1][ki-1]))
				}
			}
		}
		c.k = k
	}
	if n > c.n {
		for ni := c.n + 1; ni < n+1; ni++ {
			// Append 1
			c.values = append(c.values, []*big.Int{c.values[0][0]})
			for ki := 1; ki < c.k+1; ki++ {
				c.values[ni] = append(
					c.values[ni],
					new(big.Int).Add(c.values[ni][ki-1], c.values[ni-1][ki-1]))
			}
		}
		c.n = n
	}
	return result.Set(c.values[n][k])
}
