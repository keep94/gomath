package gomath

import (
  "math/big"
)

const (
  kNoMoreValues = "No more values on channel"
  kNotGreater = "n is not greater than the number of values already taken from channel"
)

// BigIntChan wraps a <-chan *big.Int and is used to get the Nth value off
// the wrapped channel via the Nth method. Values passed to Nth must be in
// increasing order. BigIntChan instances are not safe to use with multiple
// goroutines.
type BigIntChan struct {
  ch <-chan *big.Int
  numTaken int64
}

// NewBigIntChan creates a new instance that wraps ch.
func NewBigIntChan(ch <-chan *big.Int) *BigIntChan {
  return &BigIntChan{ch: ch}
}

// Nth returns the nth big.Int taken from the channel.
// Nth panics if n is not greater than the number of values already taken from
// the channel, or if n is greater than the total number of values in the
// channel,
func (b *BigIntChan) Nth(n int64) *big.Int {
  result, ok := b.safeNth(n)
  if !ok {
    panic(kNoMoreValues)
  }
  return result
}

func (b *BigIntChan) safeNth(n int64) (result *big.Int, ok bool) {
  if n <= b.numTaken {
    panic(kNotGreater)
  }
  for n > b.numTaken {
    result, ok = <-b.ch
    if !ok {
      return
    }
    b.numTaken++
  }
  return
}

// IntChan wraps a <-chan int64 and is used to get the Nth value off
// the wrapped channel via the Nth or SafeNth methods. Values passed to
// Nth or SafeNth must be in increasing order. IntChan instances are not
// safe to use with multiple goroutines.
type IntChan struct {
  ch <-chan int64
  numTaken int64
}

// NewIntChan creates a new instance that wraps ch
func NewIntChan(ch <-chan int64) *IntChan {
  return &IntChan{ch: ch}
}

// Nth returns the nth int64 taken from the channel.
// Nth panics if n is not greater than the number of values already taken from
// the channel, or if n is greater than the total number of values in the
// channel,
func (i *IntChan) Nth(n int64) int64 {
  result, ok := i.SafeNth(n)
  if !ok {
    panic(kNoMoreValues)
  }
  return result
}

// SafeNth works like Nth except that instead of panicing if n is greather
// than the total number of values in the channel, it returns ok=false.
func (i *IntChan) SafeNth(n int64) (result int64, ok bool) {
  if n <= i.numTaken {
    panic(kNotGreater)
  }
  for n > i.numTaken {
    result, ok = <-i.ch
    if !ok {
      return
    }
    i.numTaken++
  }
  return
}
