package factor

func PrimeFactorsByTrial(n int64) []int64 {
    factors := []int64{}
    var d int64 = 2
    var lastAppended int64 = 0
    for ; n > 1 ; {
        for ; n % d == 0 ; {
            if lastAppended != d {
                factors = append(factors, d)
                lastAppended = d
            }
            n = n / d
        }
        d = d + 1
        if d * d > n {
            if n > 1 {
                factors = append(factors, n)
                break
            }
        }
    }
    return factors
}
