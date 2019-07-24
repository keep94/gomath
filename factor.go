package gomath

// PrimePower represents a single term of a prime power decomposition
type PrimePower struct {
  Prime int64
  Power int
}

// Factor returns the prime power decomposition of n
func Factor(n int64) []PrimePower {
  if n < 1 {
    panic("n can't be less than 1")
  }
  fact := int64(2)
  var result []PrimePower
  for fact <= n / fact {
    exp := 0
    for n % fact == 0 {
      n /= fact
      exp++
    }
    if exp > 0 {
      result = append(result, PrimePower{Prime: fact, Power: exp})
    }
    fact++
  }
  if n > 1 {
    result = append(result, PrimePower{Prime: n, Power: 1})
  }
  return result
}
