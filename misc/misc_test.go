package misc

import "testing"

func TestGoldbacksOtherConjecture(t *testing.T) {
    primes := []int64{2, 3, 5, 7}
    gb := GoldbacksOtherConjecture(9, primes)
    if gb.prime != 7 {
        t.Error("Wrong Prime for GoldbacksOtherConjecture")
    }
    if gb.square != 1 {
        t.Error("Wrong Square for GoldbacksOtherConjecture")
    }
}
