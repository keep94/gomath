package gomath

import (
	"math"
	"math/big"
)

const (
	kNoMoreValues          = "No more values on stream"
	kNotGreater            = "n is not greater than the number of values already taken from stream"
	kValueLessThanLastCall = "value parameter decreased"
)

// BigIntStream represents an infinite stream of big.Int values.
type BigIntStream interface {
	// Next stores the next big.Int value at value and returns value.
	Next(value *big.Int) *big.Int
}

// IntStream represents a finite stream of int64 values.
type IntStream interface {
	// Next returns the next int64 value. If there are no more values, Next
	// returns 0, false.
	Next() (int64, bool)
}

// NthBigInt wraps a BigIntStream and provides the Nth method to return the
// Nth big.Int off the stream.
type NthBigInt struct {
	stream   BigIntStream
	numTaken int64
}

// NewNthBigInt creates a new instance that wraps stream.
func NewNthBigInt(stream BigIntStream) *NthBigInt {
	return &NthBigInt{stream: stream}
}

// Nth returns the nth big.Int taken from the wrapped stream. n is 1-based so
// 1 returns the first big.Int, 2 returns the second etc. The nth big.Int is
// stored at value, and Nth returns this value.
// Nth panics if n is not greater than the number of values already taken
// from the stream.
func (b *NthBigInt) Nth(n int64, value *big.Int) *big.Int {
	if n <= b.numTaken {
		panic(kNotGreater)
	}
	for n > b.numTaken {
		b.stream.Next(value)
		b.numTaken++
	}
	return value
}

// NthInt wraps an IntStream and provides the Nth method to return the
// Nth int64 off the stream.
type NthInt struct {
	stream   IntStream
	numTaken int64
}

// NewNthInt creates a new instance that wraps stream.
func NewNthInt(stream IntStream) *NthInt {
	return &NthInt{stream: stream}
}

// Nth returns the nth int64 taken from the wrapped stream. n is 1-based so
// 1 returns the first int64, 2 returns the second etc.
// Nth panics if n is not greater than the number of values already taken
// from the stream or if there are fewer than N total values on the stream.
func (i *NthInt) Nth(n int64) int64 {
	result, ok := i.SafeNth(n)
	if !ok {
		panic(kNoMoreValues)
	}
	return result
}

// SafeNth works like Nth except that instead of panicing if n is greather
// than the total number of values in the stream, it returns ok=false.
func (i *NthInt) SafeNth(n int64) (result int64, ok bool) {
	if n <= i.numTaken {
		panic(kNotGreater)
	}
	for n > i.numTaken {
		result, ok = i.stream.Next()
		if !ok {
			return
		}
		i.numTaken++
	}
	return
}

// BigIntCounter counts how many values on wrapped BigIntStream are less than
// or equal to a given value. Wrapped stream must be monotone increasing.
type BigIntCounter struct {
	stream    BigIntStream
	numTaken  int64
	lastTaken *big.Int
	lastCall  *big.Int
}

// NewBigIntCounter returns a BigIntCounter that wraps stream.
func NewBigIntCounter(stream BigIntStream) *BigIntCounter {
	result := &BigIntCounter{
		stream:    stream,
		lastTaken: new(big.Int),
	}
	result.stream.Next(result.lastTaken)
	result.numTaken++
	return result
}

// CountLE returns how many values on wrapped stream are less than or equal to
// value. CountLE panics if value is less than value in the previous
// call to CountLE.
func (b *BigIntCounter) CountLE(value *big.Int) int64 {
	if b.lastCall != nil && value.Cmp(b.lastCall) < 0 {
		panic(kValueLessThanLastCall)
	}
	if b.lastCall == nil {
		b.lastCall = new(big.Int)
	}
	b.lastCall.Set(value)
	for b.lastTaken.Cmp(value) < 0 {
		b.stream.Next(b.lastTaken)
		b.numTaken++
	}
	if b.lastTaken.Cmp(value) == 0 {
		return b.numTaken
	}
	return b.numTaken - 1
}

// IntCounter counts how many values on wrapped IntStream are less than or
// equal to a given value. Wrapped stream must be monotone increasing.
type IntCounter struct {
	stream    IntStream
	numTaken  int64
	lastTaken int64
	lastCall  int64
}

// NewIntCounter returns an IntCounter that wraps stream.
func NewIntCounter(stream IntStream) *IntCounter {
	return &IntCounter{
		stream:    stream,
		lastTaken: math.MinInt64,
		lastCall:  math.MinInt64,
	}
}

// CountLE returns how many values on wrapped stream are less than or equal to
// value. CountLE panics if value is less than value in the previous
// call to CountLE.
func (i *IntCounter) CountLE(value int64) int64 {
	if value < i.lastCall {
		panic(kValueLessThanLastCall)
	}
	i.lastCall = value
	for i.lastTaken < value {
		nextTaken, ok := i.stream.Next()
		if !ok {
			break
		}
		i.lastTaken = nextTaken
		i.numTaken++
	}
	if i.lastTaken <= value {
		return i.numTaken
	}
	return i.numTaken - 1
}
