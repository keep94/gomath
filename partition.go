package gomath

import (
  "math/big"
)

// Partitions generates p(1), p(2), p(3), ... where p is the partition function.
// See Partition.
func Partitions() BigIntStream {
  return &partitionStream{current: 1, partition: NewPartition()}
}

// Partition computes the partition function, p, which calculates how many ways
// n can be partitioned when order doesn't matter.
//
// For example: p(4) = 5 because there are 5 ways to express 4 as a sum
// when order doesn't matter.
// 4 = 1+1+1+1,
// 4 = 1+1+2 (covers 2+1+1 and 1+2+1),
// 4 = 1+3 (covers 3+1),
// 4 = 2+2,
// 4 = 4
//
// Partition instances are not safe to use with multiple goroutines.
type Partition struct {
  values []*big.Int
}

// NewPartition creates a new Partition instance
func NewPartition() *Partition {
  values := []*big.Int{big.NewInt(1)}
  return &Partition{values: values}
}

// Eval evaluates p(n). Eval stores the result in result and returns result.
// Eval panics if n < 0.
func (p *Partition) Eval(n int, result *big.Int) *big.Int {
  if n < 0 {
    panic("n must be greater than or equal to 0")
  }
  for n >= len(p.values) {
    p.evalNext()
  }
  return result.Set(p.values[n])
}

// Chart is used to make a chart of the partition function using the
// github.com/keep94/gochart package.
// p.Chart(n, result) is the same as p.Eval(int(n), result)
func (p *Partition) Chart(n int64, result *big.Int) *big.Int {
  return p.Eval(int(n), result)
}

func (p *Partition) evalNext() {
  idx := len(p.values)
  bigDec := 1
  smallDec := 1
  sum := new(big.Int)
  for idx >= 0 {
    f := sum.Add
    if smallDec % 2 == 0 {
      f = sum.Sub
    }
    idx -= bigDec
    if idx >= 0 {
      f(sum, p.values[idx])
    }
    idx -= smallDec
    if idx >= 0 {
      f(sum, p.values[idx])
    }
    bigDec += 2
    smallDec += 1
  }
  p.values = append(p.values, sum)
}

type partitionStream struct {
  current int
  partition *Partition
}

func (p *partitionStream) Next(value *big.Int) *big.Int {
  if value != nil {
    p.partition.Eval(p.current, value)
  }
  p.current++
  return value
}
