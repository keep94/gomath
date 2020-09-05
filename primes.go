package gomath

import (
	"math"
)

const (
	kPrimeMinSieveSize = 1000
	kPrimeMaxSieveSize = 1000000
)

// Primes generates the prime numbers in order that are greater than or
// equal to start.
func Primes(start int64) IntStream {
	if start <= 2 {
		start = 2
	} else {
		start = start/2*2 + 1 // start on odd
	}
	return &primeStream{start: start}
}

// DecadePrimes generates all x >= start in ascending order such that
// 10x + 1, 10x + 3, 10x + 7, and 10x + 9 are all prime. DecadePrimes stops
// generating when 10*x > math.MaxInt64. If start*10 > math.MaxInt64,
// DecadePrimes generates nothing.
func DecadePrimes(start int64) IntStream {
	if start < 1 {
		start = 1
	}
	if start > math.MaxInt64/10 {
		start = math.MaxInt64 / 10
	}
	return &decadeStream{primes: Primes(10*start + 1)}
}

type decadeStream struct {
	primes     IntStream
	lastDecade int64
	count      int
}

func (d *decadeStream) Next() (result int64, ok bool) {
	for {
		prime, pok := d.primes.Next()
		if !pok {
			return
		}
		decade := prime / 10
		if decade == d.lastDecade {
			d.count++
			if d.count == 4 {
				result = decade
				ok = true
				return
			}
		} else {
			d.lastDecade = decade
			d.count = 1
		}
	}
}

type primeStream struct {
	start int64
	idx   int
	sieve []bool
}

func newSize(oldSize int) int {
	result := oldSize * 2
	if result < kPrimeMinSieveSize {
		result = kPrimeMinSieveSize
	}
	if result > kPrimeMaxSieveSize {
		result = kPrimeMaxSieveSize
	}
	return result
}

func divideRoundUpOdd(n, d int64) int64 {
	roundUp := (n-1)/d + 1
	return roundUp/2*2 + 1
}

func (p *primeStream) fillInSieve() {
	end := p.start + 2*int64(len(p.sieve))
	for factor := int64(3); factor <= end/factor; factor += 2 {
		multStart := divideRoundUpOdd(p.start, factor)
		if multStart < factor {
			multStart = factor
		}
		multEnd := divideRoundUpOdd(end, factor)
		for i := multStart; i < multEnd; i += 2 {
			tableOffset := (factor*i - p.start) / 2
			p.sieve[tableOffset] = true
		}
	}
}

func (p *primeStream) currentValue() int64 {
	return p.start + 2*int64(p.idx)
}

func (p *primeStream) Next() (result int64, ok bool) {
	if p.start == 2 {
		result = 2
		ok = true
		p.start++
		return
	}
	for {
		if p.idx == len(p.sieve) {
			p.start = p.currentValue()
			p.idx = 0
			nextSize := int64(newSize(len(p.sieve)))
			p.sieve = nil
			maxSize := (math.MaxInt64 - p.start) / 2
			if maxSize == 0 {
				return
			}
			if nextSize > maxSize {
				nextSize = maxSize
			}
			p.sieve = make([]bool, nextSize)
			p.fillInSieve()
		}
		if !p.sieve[p.idx] {
			result = p.currentValue()
			ok = true
			p.idx++
			return
		}
		p.idx++
	}
}
