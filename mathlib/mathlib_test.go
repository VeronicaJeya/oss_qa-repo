package mathlib

import "testing"

// Fast sanity test
func TestBasicMath(t *testing.T) {
    if Add(5, 5) != 10 {
        t.Fail()
    }
}

// Long-running tests to INCREASE Jenkins build time
func TestHeavyPrimeCount_Long(t *testing.T) {
    // Bigger limit -> more time
    result := HeavyPrimeCount(200000)
    if result <= 0 {
        t.Fail()
    }
}

func TestBigFactorial_Long(t *testing.T) {
    // Much larger factorial -> more time
    res := BigFactorial(25000)
    if res == nil {
        t.Fail()
    }
}

func TestTrigSeries_Long(t *testing.T) {
    // More iterations -> more time
    res := TrigSeries(150_000_000)
    if res == 0 {
        t.Fail()
    }
}
