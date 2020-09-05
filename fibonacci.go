package gomath

import (
	"math/big"
)

// Fibonacci generates fibonacci numbers. first and second are the
// first and second terms in the sequence, normally 1 and 1.
func Fibonacci(first, second int64) BigIntStream {
	return &fibStream{a: big.NewInt(first), b: big.NewInt(second)}
}

type fibStream struct {
	a *big.Int
	b *big.Int
}

func (f *fibStream) Next(value *big.Int) *big.Int {
	value.Set(f.a)
	f.a.Add(f.a, f.b)
	f.a, f.b = f.b, f.a
	return value
}
