package misc

type GoldbachForm struct {
    prime int64
    square int64
}

func GoldbacksOtherConjecture(n int64, primes []int64) GoldbachForm {
    i := 0
    var j int64 = 0
    for ; n >= primes[i] ; i = i + 1 {
        for j = 0; (2 * j * j) <= (n - primes[i]) ; j = j + 1 {
            if primes[i] + 2 * j * j == n {
                return GoldbachForm{primes[i], j}
            }
        }
    }
    return GoldbachForm{0, 0}
}
