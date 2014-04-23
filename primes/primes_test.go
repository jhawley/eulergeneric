package primes

import "testing"

func TestSieveOfEratosthenes(t *testing.T) {
    primes := SieveOfEratosthenes(10)
    if len(primes) == 4 && primes[3] != 7 {
        t.Error("SieveOfEratosthenes failed to return primes correctly")
    }
    primes = SieveOfEratosthenes(11)
    if len(primes) == 5 && primes[4] != 11 {
        t.Error("SieveOfEratosthenes failed to return primes correctly")
    }
    primes = SieveOfEratosthenes(10000)
}

func TestSieveOfEratosthenesSequential(t *testing.T) {
    primes := SieveOfEratosthenes(10)
    if len(primes) == 4 && primes[3] != 7 {
        t.Error("SieveOfEratosthenesSequential failed to return primes correctly")
    }
    primes = SieveOfEratosthenes(11)
    if len(primes) == 5 && primes[4] != 11 {
        t.Error("SieveOfEratosthenesSequential failed to return primes correctly")
    }
    primes = SieveOfEratosthenes(10000)
}

func BenchmarkSieveOfEratosthenes100000(b *testing.B) {
    for i := 0 ; i < b.N; i++ {
        SieveOfEratosthenes(100000)
    }
}

func BenchmarkSieveOfEratosthenesSequential100000(b *testing.B) {
    for i := 0 ; i < b.N; i++ {
        SieveOfEratosthenesSequential(100000)
    }
}
