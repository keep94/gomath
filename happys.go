package gomath

import (
  "context"
  "math"
)

// Happys generates all the happy numbers in order that are greater than or
// equal to start.
func Happys(ctx context.Context, start int64) <-chan int64 {
  if start < 1 {
    start = 1
  }
  result := make(chan int64)
  go func() {
    defer close(result)
    he := newHappyEngine()
    for {
      if he.isHappy(start) {
        select {
          case <-ctx.Done():
            return
          case result <- start:
        }
      }
      if start == math.MaxInt64 {
        return
      }
      start++
    }
  }()
  return result
}

type happyStatus uint8

const (
  unknown happyStatus = iota
  unhappy
  happy
)

type happyEngine struct {
  table [1000]happyStatus
}

func newHappyEngine() *happyEngine {
  var result happyEngine
  result.table[1] = happy
  return &result
}

func (h *happyEngine) isHappy(x int64) bool {
  if x >= int64(len(h.table)) {
    return h.isHappy(sumSquareDigits(x))
  }
  switch h.table[x] {
    case happy:
      return true
    case unhappy:
      return false
    case unknown:
      h.table[x] = unhappy
      result := h.isHappy(sumSquareDigits(x))
      if result {
        h.table[x] = happy
      }
      return result
    default:
      panic("Bad value in h.table")
  }
}

func sumSquareDigits(x int64) int64 {
  result := int64(0)
  for x != 0 {
    dig := x % 10
    result += dig*dig
    x /= 10
  }
  return result
}
