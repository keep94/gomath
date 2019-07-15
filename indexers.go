package gomath

import (
  "math/big"
)

const (
  kNLessThan1 = "n was less than 1"
  kNoMoreValues = "No more values on channel"
)

// BigIntIndexer is used to find the nth big.Int taken from a channel.
// It does this by storing all the values it takes from the channel.
// BigIntIndexer instances are not safe to use with multiple goroutines.
// BigIntIndexer instances panic if their channel runs out of values or if
// they encounter nil values on their channel.
type BigIntIndexer struct {
  ch <-chan *big.Int
  values []*big.Int
}

// NewBigIntIndexer creates a new instance from a channel
func NewBigIntIndexer(ch <-chan *big.Int) *BigIntIndexer {
  return &BigIntIndexer{ch: ch}
}

// Nth stores in result the nth big.Int taken from the channel consuming the
// channel as needed and returns result. Nth panics if n < 1
func (b *BigIntIndexer) Nth(n int, result *big.Int) *big.Int {
  if n < 1 {
    panic(kNLessThan1)
  }
  for len(b.values) < n {
    val, cok  := <-b.ch
    if !cok {
      panic(kNoMoreValues)
    }
    if val == nil {
      panic("nil *big.Int encountered on channel")
    }
    b.values = append(b.values, val)
  }
  return result.Set(b.values[n-1])
}

// IntIndexer is used to find the nth int64 taken from a channel.
// It does this by storing all the values it takes from the channel.
// IntIndexer instances are not safe to use with multiple goroutines.
// IntIndexer instances panic if their channel runs out of values.
type IntIndexer struct {
  ch <-chan int64
  values []int64
}

// NewIntIndexer creates a new instance from a channel
func NewIntIndexer(ch <-chan int64) *IntIndexer {
  return &IntIndexer{ch: ch}
}

// Nth returns the nth int64 taken from the channel consuming the
// channel as needed. Nth panics if n < 1.
func (i *IntIndexer) Nth(n int) int64 {
  if n < 1 {
    panic(kNLessThan1)
  }
  for len(i.values) < n {
    val, cok  := <-i.ch
    if !cok {
      panic(kNoMoreValues)
    }
    i.values = append(i.values, val)
  }
  return i.values[n-1]
}
