package gomath

import (
  "math"
  "math/big"
)

const (
  kNoMoreValues = "No more values on channel"
  kNotGreater = "n is not greater than the number of values already taken from channel"
  kValueLessThanLastCall = "value parameter decreased"
)

// BigIntChan wraps a <-chan *big.Int and is used to get the Nth value off
// the wrapped channel via the Nth method. Values for n passed to Nth must
// be in increasing order. BigIntChan instances are not safe to use with
// multiple goroutines.
type BigIntChan struct {
  ch <-chan *big.Int
  numTaken int64
}

// NewBigIntChan creates a new instance that wraps ch.
func NewBigIntChan(ch <-chan *big.Int) *BigIntChan {
  return &BigIntChan{ch: ch}
}

// Nth returns the nth big.Int taken from the wrapped channel.
// Nth panics if n is not greater than the number of values already taken
// from the channel or if n is greater than the total number of values in
// the channel,
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
// the wrapped channel via the Nth method. Values for n passed to Nth must
// be in increasing order. IntChan instances are not safe to use with
// multiple goroutines.
type IntChan struct {
  ch <-chan int64
  numTaken int64
}

// NewIntChan creates a new instance that wraps ch
func NewIntChan(ch <-chan int64) *IntChan {
  return &IntChan{ch: ch}
}

// Nth returns the nth int64 taken from the wrapped channel.
// Nth panics if n is not greater than the number of values already taken
// from the channel or if n is greater than the total number of values in
// the channel,
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

// BigIntCounter counts how many f(1), f(2), f(3),... are less than a given
// value. f must be monotone increasing. BigIntCounter will evaluate f(x)
// at most once. If y < x, BigIntCounter evaluates f(y) before it evaluates
// f(x).
type BigIntCounter struct {
  f func(int64) *big.Int
  numTaken int64
  lastTaken *big.Int
  lastCall *big.Int
}

// NewBigIntCounter returns a BigIntCounter that wraps f.
func NewBigIntCounter(f func(int64) *big.Int) *BigIntCounter {
  return &BigIntCounter{
      f: f,
      numTaken: 1,
      lastTaken: f(1),
  }
}

// CountLE returns how many f(1), f(2), f(3), ... are less than or equal to
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
    b.numTaken++
    b.lastTaken = b.f(b.numTaken)
  }
  if b.lastTaken.Cmp(value) == 0 {
    return b.numTaken
  }
  return b.numTaken - 1
}

// IntCounter counts how many f(1), f(2), f(3),... are less than a given
// value. f must be monotone increasing. IntCounter will evaluate f(x)
// at most once. If y < x, IntCounter evaluates f(y) before it evaluates
// f(x).
type IntCounter struct {
  f func(int64) int64
  numTaken int64
  lastTaken int64
  lastCall int64
}

// NewIntCounter returns an IntCounter that wraps f.
func NewIntCounter(f func(int64) int64) *IntCounter {
  return &IntCounter{
      f: f,
      numTaken: 1,
      lastTaken: f(1),
      lastCall: math.MinInt64,
  }
}

// CountLE returns how many f(1), f(2), f(3), ... are less than or equal to
// value. CountLE panics if value is less than value in the previous
// call to CountLE.
func (i *IntCounter) CountLE(value int64) int64 {
  if value < i.lastCall {
    panic(kValueLessThanLastCall)
  }
  i.lastCall = value
  for i.lastTaken < value {
    i.numTaken++
    i.lastTaken = i.f(i.numTaken)
  }
  if i.lastTaken == value {
    return i.numTaken
  }
  return i.numTaken - 1
}
