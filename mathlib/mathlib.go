
package mathlib

import (
    "math"
    "sort"
    "math/big"
)

func Add(a, b int) int { return a + b }
func Sub(a, b int) int { return a - b }
func Mul(a, b int) int { return a * b }

func Div(a, b int) float64 {
    if b == 0 {
        return 0
    }
    return float64(a) / float64(b)
}

func Square(n float64) float64 { return n * n }
func SquareRoot(n float64) float64 { return math.Sqrt(n) }

func OrderedList(nums []int) []int {
    sort.Ints(nums)
    return nums
}

func HeavyPrimeCount(limit int) int {
    count := 0
    for n := 2; n <= limit; n++ {
        isPrime := true
        for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
            if n%i == 0 {
                isPrime = false
                break
            }
        }
        if isPrime {
            count++
        }
    }
    return count
}

func BigFactorial(n int) *big.Int {
    result := big.NewInt(1)
    for i := 1; i <= n; i++ {
        result.Mul(result, big.NewInt(int64(i)))
    }
    return result
}

func TrigSeries(iterations int) float64 {
    sum := 0.0
    for i := 1; i <= iterations; i++ {
        sum += math.Sin(float64(i)) * math.Cos(float64(i/2))
    }
    return sum
}
