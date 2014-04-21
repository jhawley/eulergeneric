package misc

import "testing"

func TestGoldbacksOtherConjecture(t *testing.T) {
    primes := []int64{2, 3, 5, 7}
    gb := GoldbacksOtherConjecture(9, primes)
    if gb.Prime != 7 {
        t.Error("Wrong Prime for GoldbacksOtherConjecture")
    }
    if gb.Square != 1 {
        t.Error("Wrong Square for GoldbacksOtherConjecture")
    }
}

func TestFilterSlice(t *testing.T) {
    f := func(x int64) bool {
        return (x % 3 == 0)
    }
    slice := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
    slice = FilterSlice(slice, f)
    if slice[2] != 9 {
        t.Error("FilterSlice failed to apply filter correctly")
    }
}
