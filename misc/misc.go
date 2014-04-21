package misc

type GoldbachForm struct {
    Prime int64
    Square int64
}

func GoldbacksOtherConjecture(n int64, primes []int64) GoldbachForm {
    i := 0
    max := len(primes)
    var j int64 = 0
    for ; (i < max) && (n >= primes[i]) ; i = i + 1 {
        for j = 0; (2 * j * j) <= (n - primes[i]) ; j = j + 1 {
            if primes[i] + 2 * j * j == n {
                return GoldbachForm{primes[i], j}
            }
        }
    }
    return GoldbachForm{0, 0}
}

func FilterSlice(s []int64, f func(int64) bool) []int64 {
    i := 0
    max := len(s)
    result := []int64{}
    for ; i < max ; i = i + 1 {
        if f(s[i]) {
            result = append(result, s[i])
        }
    }
    return result
}
