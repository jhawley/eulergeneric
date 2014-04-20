package factor

import (
    "testing"
)

func TestPrimeFactorsByTrial(t *testing.T) {
    factors := PrimeFactorsByTrial(1000)
    if len(factors) != 2 {
        t.Error("Wrong number of factors")
    }
    if factors[0] != 2 {
        t.Error("Wrong factors")
    }
}
