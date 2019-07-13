package gomath_test

import (
  "math"
  "math/big"
  "testing"
)


func assertBigIntEqual(t *testing.T, expected int64, actual *big.Int) {
  t.Helper()
  if actual.Cmp(big.NewInt(expected)) != 0 {
    t.Errorf("Expected %v, actual %v", expected, actual)
  }
}

func assertBigIntNil(t *testing.T, actual *big.Int) {
  t.Helper()
  if actual != nil {
    t.Error("Expected nil.")
  }
}

func assertFalse(t *testing.T, ok bool) {
  t.Helper()
  if ok {
    t.Error("Expected false")
  }
}

func assertTrue(t *testing.T, ok bool) {
  t.Helper()
  if !ok {
    t.Error("Expected true")
  }
}

func assertCloseTo(t *testing.T, expected float64, actual float64) {
  t.Helper()
  if math.Abs((expected - actual) / expected) > 0.0001 {
    t.Errorf("Expected %v, got %v", expected, actual)
  }
}

func assertPanic(t *testing.T, f func()) {
  t.Helper()
  defer func() {
    recover()
  }()
  f()
  t.Error("Expected panic")
}

func assertEqual(
    t *testing.T, expected, actual interface{}) {
  t.Helper()
  if expected != actual {
    t.Errorf("Expected %v, got %v", expected, actual)
  }
}

func checkInfInt64Chan(
    t *testing.T, ch <-chan int64, expectedValues ...int64) {
  t.Helper()
  for _, expected := range expectedValues {
    actual, ok := <-ch
    if !ok {
        t.Fatal("No more values on channel")
    }
    if actual != expected {
      t.Fatalf("Expected %v, got %v", expected, actual)
    }
  }
}
