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

func TestNumberSpiralDiagonals(t *testing.T) {
    ds := NumberSpiralDiagonals(3)
    if ds[8] != 25 {
        t.Error("NumberSpiralDiagonals missed last number")
    }
    var sum int64 = 0
    i := 0
    for ; i < len(ds) ; i = i + 1 {
        sum = sum + ds[i]
    }
    if sum != 101 {
        t.Error("NumberSpiralDiagonals has the wrong elements")
    }
}

func TestIsPandigital(t *testing.T) {
    if IsPandigital(123456789) {
        t.Error("IsPandigital incorrectly verified a simple pandigital")
    }
    if !IsPandigital(192) {
        t.Error("IsPandigital failed to verify a complex pandigital")
    }
    if IsPandigital(123456788) {
        t.Error("IsPandigital incorrectly verified a nonpandigital")
    }
    if IsPandigital(193) {
        t.Error("IsPandigital incorrectly verified a complex nonpandigital")
    }
}
