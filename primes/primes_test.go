package primes

import "testing"

func TestSieveOfEratosthenes(t *testing.T) {
    primes := SieveOfEratosthenes(10)
    if len(primes) == 4 && primes[3] != 7 {
        t.Error("SieveOfEratosthenes failed to return primes correctly")
    }
}
