package misc

import  (
    "strconv"
    "strings"
    "math/big"
)

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

func NumberSpiralDiagonals(squareId int64) []int64 {
    var i int64 = 2
    j := 0
    spiral := []int64{1}
    var last int64 = 1
    for ; i <= squareId ; i = i + 1 {
        for j = 0; j < 4; j = j + 1 {
            last = last + 2 * (i - 1)
            spiral = append(spiral, last)            
        }
    }
    return spiral
}

func IsPandigital(n int64) bool {
    items := []int64{n}
    var multiplier int64 = 2
    for ; ; multiplier = multiplier + 1 {
        switch checkPandigital(items) {
        case -1:
            return false
        case 1:
            return true
        default:
            items = append(items, n * multiplier)
        }
    }
    return false
}

// -1 if chars > 9 or = 9 and not pandigital, 0 if not pandigital, 1 otherwise
func checkPandigital(ns []int64) int {
    strs := []string{}
    i := 0
    max := len(ns)
    for ; i < max ; i = i + 1 {
        strs = append(strs, strconv.FormatInt(ns[i], 10))
    }
    complete := strings.Join(strs, "")
    reader := strings.NewReader(complete)
    if(reader.Len() > 9) {
        return -1
    }
    if(reader.Len() < 9) {
        return 0
    }
    if(strings.Contains(complete, "1") && strings.Contains(complete, "2") &&
      strings.Contains(complete, "3") && strings.Contains(complete, "4") &&
      strings.Contains(complete, "5") && strings.Contains(complete, "6") &&
      strings.Contains(complete, "7") && strings.Contains(complete, "8") &&
      strings.Contains(complete, "9") && len(ns) > 1) {
        return 1
    } else {
        return -1
    }
    return -1
}

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func ReverseBigInt(n *big.Int) *big.Int {
    g := big.NewInt(0)
    g.SetString(reverse(n.String()), 10)
    return g
}

func IsPalindrome(n *big.Int) bool {
    g := ReverseBigInt(n)
    return g.String() == n.String()
}

func IsLychrel(n *big.Int, testsRun int) bool {
    g := ReverseBigInt(n)
    result := big.NewInt(0)
    result.Add(g, n)
    if IsPalindrome(result) {
        return false
    } else if testsRun == 48 {
        return true
    } else {
        testsRun++
        return IsLychrel(result, testsRun)
    }
    return false
}
