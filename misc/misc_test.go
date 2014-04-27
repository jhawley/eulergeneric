package misc

import (
    "testing"
    "math/big"
)

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

func TestReverseBig(t *testing.T) {
    n := big.NewInt(1234)
    g := ReverseBigInt(n)
    if g.String() != "4321" {
        t.Error("ReverseBigInt failed to reverse big.Int")
    }
}

func TestIsPalindrome(t *testing.T) {
    n := big.NewInt(7337)
    if !IsPalindrome(n) {
        t.Error("IsPalindrome failed to detect palindrome")
    }
    m := big.NewInt(7333)
    if IsPalindrome(m) {
        t.Error("IsPalindrome failed to reject non-palindrome")
    }
}

func TestIsLychrel(t *testing.T) {
    n1 := big.NewInt(47)
    n2 := big.NewInt(349)
    n3 := big.NewInt(196)
    if IsLychrel(n1, 0) {
        t.Error("IsLychrel incorrectly identified a simple non-Lychrel")
    }
    if IsLychrel(n2, 0) {
        t.Error("IsLychrel incorrectly identified a complex non-Lychrel")
    }
    if !IsLychrel(n3, 0) {
        t.Error("IsLychrel incorrectly identified a Lychrel")
    }
}
