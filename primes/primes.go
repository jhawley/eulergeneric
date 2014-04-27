package primes

import "github.com/jhawley/eulergeneric/misc"

func SieveOfEratosthenes(limit int64) []int64 {
    var primes = make(map[int64]bool)
    result := []int64{}
    c := make(chan []int64)
    var i int64 = 3
    var j int64 = 0
    primes[2] = true
    for ; i <= limit ; i++ {
        primes[i] = true
    }
    go buildSieve(2, limit, c)
    for i = 3; i <= limit ; i = i + 2 {
        go buildSieve(i, limit, c)
    }
    temp := []int64{}
    for i = 0; i < (((limit - 3)/2) + 2); i++ {
        temp = <-c
        for j = 0; j < int64(len(temp)); j++ {
            primes[temp[j]] = false            
        }
    }
    for i = 0; i <= limit; i++ {
        if(primes[i]) {
            result = append(result, i)
        }
    }
    return result
}

func buildSieve(n int64, limit int64, c chan []int64) {
    var i int64 = n
    multiples := []int64{}
    for ; i <= limit ; i = i + 1 {
        if !((n == i) || (i % n != 0)) {
            multiples = append(multiples, i)
        }
    }
    c <- multiples
}

func SieveOfEratosthenesSequential(limit int64) []int64 {
    primes := []int64{2}
    var i int64 = 3
    for ; i <= limit ; i = i + 2 {
        primes = append(primes, i)
    }
    var j int64 = 3
    for ; j <= limit ; j = j + 2 {
        primes = misc.FilterSlice(primes, func(x int64) bool {
            return ((x == j) || (x % j != 0))
        })
    }
    return primes
}
