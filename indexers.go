package gomath

import (
  "math/big"
)

// BigIntIndexer is used to find the nth big.Int taken from a channel.
// It does this by tracking all the values it takes from the channel.
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
    b.values = append(b.values, val)
  }
  return result.Set(b.values[n-1]), true
}
