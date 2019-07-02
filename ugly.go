package gomath

import (
  "context"
  "container/heap"
  "math/big"
)

// Ugly returns the numbers whose prime factors are a subset of primeFactors
// in ascending order.
func Ugly(ctx context.Context, primeFactors ...int64) <-chan *big.Int {
  checkPrimeFactors(primeFactors)
  result := make(chan *big.Int)
  valueEntryTail := &valueEntry{value: big.NewInt(1)}
  var hp factorPointerHeap
  createFactorPointersOnHeap(primeFactors, valueEntryTail, &hp)
  go func() {
    defer close(result)
    for {
      select {
        case <-ctx.Done():
          return
        case result <- new(big.Int).Set(valueEntryTail.value):
      }
      fp := heap.Pop(&hp).(*factorPointer)
      for fp.effectiveValue.Cmp(valueEntryTail.value) == 0 {
        fp.advance()
        heap.Push(&hp, fp)
        fp = heap.Pop(&hp).(*factorPointer)
      }
      valueEntryTail = fp.appendEffectiveValue(valueEntryTail)
      fp.advance()
      heap.Push(&hp, fp)
    }
  }()
  return result
}

// Ugly works by keeping a linked list of valueEntry instances arranged in 
// ascending order by value. The valueEntry instances represent the values
// that Ugly has already returned. For each prime factor passed to Ugly, Ugly
// keeps a pointer to the valueEntry instance where that value times the
// prime factor is smallest value equal to or greater than the last value
// Ugly returned. These pointers are called factorPointers, and the value that
// the factorPointer points to times the prime factor is called the
// effectiveValue of the factor pointer.  For example, if the prime factors
// are 2, 3, 5, and Ugly just returned is 500, then the
// 2 factorPointer would point to 250 and have an effectiveValue of 500,
// the 3 factorPointer would point to 180, the smallest value just greater
// than 166 2/3,  and have an effective value of 540, and the 5 factor
// pointer would point to 100 and have an effective value of 500.
//
// To find the next value to return, Ugly uses the factorPointer
// with the smallest effectiveValue and advances that factorPointer so that
// it points to its next value. Note that Ugly won't return duplicate values.
// So in our example above, to determine the next value to return, Ugly would
// take the 2 factorPointer because it has the smallest effectiveValue of 500,
// but it skips returning 500 because of duplicates, and advances the 2
// factorPointer so that it points to 256 and has an effectiveValue of 512.
// Now the 5 factorPointer has the smallest effectiveValue of 500, so Ugly
// uses that, but it skips returning 500 because of duplicates, and advances
// the 5 factorPointer so that it points to 108 and has an effectiveValue
// of 540.  Now the 2 factorPointer has the smallest effectiveValue of 512,
// because the 3 and 5 factorPointers have an effectiveValue of 540. 
// Ugly adds 512 to the end of the singly linked list of valueEntry
// instances and returns it. Then it advances the 2 factor pointer.

// Note that as Ugly returns values, the factorPointers are moved up the list
// of valueEntries. Note that once all the factorPointers have moved passed
// a valueEntry instance then that instance can be garbage collected.

// valueEntry instances form a singly linked list of values that Ugly has
// already returned in ascending order.
type valueEntry struct {
  value *big.Int
  next *valueEntry
}

// factorPointer represents the factorPointer instances mentioned above
type factorPointer struct {
  primeFactor *big.Int
  valuePtr *valueEntry
  effectiveValue *big.Int
}

func newFactorPointer(
    primeFactor int64, valuePtr *valueEntry) *factorPointer {
  result := &factorPointer{
      primeFactor: big.NewInt(primeFactor),
      valuePtr: valuePtr}
  result.effectiveValue = new(big.Int).Mul(
      result.primeFactor, result.valuePtr.value)
  return result
}

func (f *factorPointer) advance() {
  f.valuePtr = f.valuePtr.next
  f.effectiveValue = f.effectiveValue.Mul(f.primeFactor, f.valuePtr.value)
}

func (f *factorPointer) appendEffectiveValue(tail *valueEntry) *valueEntry {
  result := new(valueEntry)
  result.value = new(big.Int).Set(f.effectiveValue)
  tail.next = result
  return result
}

type factorPointerHeap []*factorPointer

func (h factorPointerHeap) Len() int {
  return len(h)
}

func (h factorPointerHeap) Less(i, j int) bool {
  return h[i].effectiveValue.Cmp(h[j].effectiveValue) < 0
}

func (h factorPointerHeap) Swap(i, j int) {
  h[i], h[j] = h[j], h[i]
}

func (h *factorPointerHeap) Push(x interface{}) {
  *h = append(*h, x.(*factorPointer))
}

func (h *factorPointerHeap) Pop() interface{} {
  old := *h
  n := len(old)
  item := old[n-1]
  *h = old[0:n-1]
  return item
}

func createFactorPointersOnHeap(
    primeFactors []int64, ptr *valueEntry, hp *factorPointerHeap) {
  for _, primeFactor := range primeFactors {
    heap.Push(hp, newFactorPointer(primeFactor, ptr))
  } 
}

func checkPrimeFactors(primeFactors []int64) {
  for _, primeFactor := range primeFactors {
    if primeFactor < 2 {
      panic("prime factors must be at least 2")
    }
  }
}
