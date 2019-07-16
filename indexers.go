package gomath

import (
  "math/big"
)

const (
  kNLessThan1 = "n was less than 1"
  kNoMoreValues = "No more values on channel"
  kNotGreater = "n is not greater than n in previous call to Nth"
  kNilBigInt = "nil *big.Int encountered on channel"
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
    val, ok  := <-b.ch
    if !ok {
      panic(kNoMoreValues)
    }
    if val == nil {
      panic(kNilBigInt)
    }
    b.values = append(b.values, val)
  }
  return result.Set(b.values[n-1])
}

// BigIntChan works like BigIntIndexer except it doesn't store values
// taken from the channel.
type BigIntChan struct {
  ch <-chan *big.Int
  numTaken int64
}

// NewBigIntChan creates a new instance from a channel
func NewBigIntChan(ch <-chan *big.Int) *BigIntChan {
  return &BigIntChan{ch: ch}
}

// Nth returns the nth big.Int taken from the channel.
// Nth panics if n < 1 or if n is not greater than the n passed
// in the previous call to Nth.
func (b *BigIntChan) Nth(n int64) *big.Int {
  if n <= b.numTaken {
    panic(kNotGreater)
  }
  var result *big.Int
  var ok bool
  for n > b.numTaken {
    result, ok = <-b.ch
    if !ok {
      panic(kNoMoreValues)
    }
    if result == nil {
      panic(kNilBigInt)
    }
    b.numTaken++
  }
  return result
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
    val, ok  := <-i.ch
    if !ok {
      panic(kNoMoreValues)
    }
    i.values = append(i.values, val)
  }
  return i.values[n-1]
}

// IntChan works like IntIndexer except it doesn't store values
// taken from the channel.
type IntChan struct {
  ch <-chan int64
  numTaken int64
}

// NewIntChan creates a new instance from a channel
func NewIntChan(ch <-chan int64) *IntChan {
  return &IntChan{ch: ch}
}

// Nth returns the nth int64 taken from the channel.
// Nth panics if n < 1 or if n is not greater than the n passed
// in the previous call to Nth.
func (i *IntChan) Nth(n int64) int64 {
  if n <= i.numTaken {
    panic(kNotGreater)
  }
  var result int64
  var ok bool
  for n > i.numTaken {
    result, ok = <-i.ch
    if !ok {
      panic(kNoMoreValues)
    }
    i.numTaken++
  }
  return result
}
