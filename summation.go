package gomath

import (
  "math/big"
)

// IntSummation generates the summation of the values in stream. For example if
// stream generates the primes, IntSummation would generate 2, 5, 10, 17, ...
func IntSummation(stream IntStream) IntStream {
  return &intSumStream{stream: stream}
}

// BigIntSummation generates the summation of the values in stream.
// For example if stream generates the primes, BigIntSummation would
// generate 2, 5, 10, 17, ...
func BigIntSummation(stream BigIntStream) BigIntStream {
  return &bigIntSumStream{
      stream: stream, nextVal: new(big.Int), sum: new(big.Int)}
}

type intSumStream struct {
  stream IntStream
  sum int64
}

func (s *intSumStream) Next() (result int64, ok bool) {
  value, vok := s.stream.Next()
  if !vok {
    return
  }
  s.sum += value
  result = s.sum
  ok = true
  return
}

type bigIntSumStream struct {
  stream BigIntStream
  nextVal *big.Int
  sum *big.Int
}

func (s *bigIntSumStream) Next(value *big.Int) *big.Int {
  s.stream.Next(s.nextVal)
  s.sum.Add(s.sum, s.nextVal)
  if value != nil {
    value.Set(s.sum)
  }
  return value
}
