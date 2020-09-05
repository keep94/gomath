package gomath_test

import (
	"math"
	"testing"

	"github.com/keep94/gomath"
)

func TestHappys(t *testing.T) {
	happys := gomath.Happys(100)
	checkInfInt64Stream(
		t,
		happys,
		100, 103, 109, 129, 130, 133, 139, 167, 176, 188)
}

func TestHappysMax(t *testing.T) {
	start := int64(math.MaxInt64 - 1000)
	happys := gomath.Happys(start)
	found := false
	happy, ok := happys.Next()
	for ; ok; happy, ok = happys.Next() {
		assertTrue(t, happy >= start)
		found = true
	}
	assertTrue(t, found)
}

func TestNthHappy(t *testing.T) {
	nth := gomath.NewNthInt(gomath.Happys(0))
	assertEqual(t, int64(100), nth.Nth(20))
	assertEqual(t, int64(694), nth.Nth(100))
	assertEqual(t, int64(6899), nth.Nth(1000))
	assertEqual(t, int64(67169), nth.Nth(10000))
}

func BenchmarkHappys(b *testing.B) {
	happys := gomath.Happys(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		happys.Next()
	}
}
