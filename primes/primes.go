package primes

import "github.com/jhawley/eulergeneric/misc"

func SieveOfEratosthenes(limit int64) []int64 {
    primes := []int64{}
    var i int64 = 2
    for ; i <= limit ; i = i + 1 {
        primes = append(primes, i)
    }
    var j int64 = 2
    for ; j <= limit ; j = j + 1 {
        primes = misc.FilterSlice(primes, func(x int64) bool {
            return ((x == j) || (x % j != 0))
        })
    }
    return primes
}
