package gomath

import (
  "math/big"
)

// BigIntIndexer is used to find the nth big.Int taken from a channel.
// It does this by tracking all the values it takes from the channel.
// BigIntIndexer instances are not safe to use with multiple goroutines.
// BigIntIndexer instances panic if they encounter nil values on their channel.
type BigIntIndexer struct {
  ch <-chan *big.Int
  values []*big.Int
}

// NewBigIntIndexer creates a new instance from a channel
func NewBigIntIndexer(ch <-chan *big.Int) *BigIntIndexer {
  return &BigIntIndexer{ch: ch}
}

// Nth stores in result the nth big.Int taken from the channel consuming the
// channel as needed and returns result, true. If n < 1 or if the channel has
// fewer values than n,  Nth returns nil, false and the value of result is
// left unchanged.
func (b *BigIntIndexer) Nth(n int, result *big.Int) (value *big.Int, ok bool) {
  if n < 1 {
    return
  }
  for len(b.values) < n {
    val, cok  := <-b.ch
    if !cok {
      return
    }
    if val == nil {
      panic("nil *big.Int encountered on channel")
    }
    b.values = append(b.values, val)
  }
  return result.Set(b.values[n-1]), true
}

// IntIndexer is used to find the nth int64 taken from a channel.
// It does this by tracking all the values it takes from the channel.
// IntIndexer instances are not safe to use with multiple goroutines.
type IntIndexer struct {
  ch <-chan int64
  values []int64
}

// NewIntIndexer creates a new instance from a channel
func NewIntIndexer(ch <-chan int64) *IntIndexer {
  return &IntIndexer{ch: ch}
}

// Nth returns the nth int64 taken from the channel consuming the
// channel as needed. If n < 1 or if the channel has fewer values than n,
// Nth returns ok=false.
func (i *IntIndexer) Nth(n int) (value int64, ok bool) {
  if n < 1 {
    return
  }
  for len(i.values) < n {
    val, cok  := <-i.ch
    if !cok {
      return
    }
    i.values = append(i.values, val)
  }
  return i.values[n-1], true
}
