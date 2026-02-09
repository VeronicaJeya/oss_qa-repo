
package main

import (
    "fmt"
    "go-math-final/mathlib"
)

func main() {
    fmt.Println("===== RUNNING FULL MATH DEMO APP =====")

    fmt.Println("\n--- Basic Math ---")
    fmt.Println("Add(10,5):", mathlib.Add(10, 5))
    fmt.Println("Sub(10,5):", mathlib.Sub(10, 5))
    fmt.Println("Mul(10,5):", mathlib.Mul(10, 5))
    fmt.Println("Div(10,5):", mathlib.Div(10, 5))
    fmt.Println("Square(7):", mathlib.Square(7))
    fmt.Println("SquareRoot(49):", mathlib.SquareRoot(49))
     fmt.Println("Add(10,5):", mathlib.Add(10, 5))
    fmt.Println("Sub(10,5):", mathlib.Sub(10, 5))
    fmt.Println("Mul(10,5):", mathlib.Mul(10, 5))
    fmt.Println("Div(10,5):", mathlib.Div(10, 5))
    fmt.Println("Square(7):", mathlib.Square(7))
    fmt.Println("SquareRoot(49):", mathlib.SquareRoot(49))
     fmt.Println("Add(10,5):", mathlib.Add(10, 5))
    fmt.Println("Sub(10,5):", mathlib.Sub(10, 5))
    fmt.Println("Mul(10,5):", mathlib.Mul(10, 5))
    fmt.Println("Div(10,5):", mathlib.Div(10, 5))
    fmt.Println("Square(7):", mathlib.Square(7))
    fmt.Println("SquareRoot(49):", mathlib.SquareRoot(49))
     fmt.Println("Add(10,5):", mathlib.Add(10, 5))
    fmt.Println("Sub(10,5):", mathlib.Sub(10, 5))
    fmt.Println("Mul(10,5):", mathlib.Mul(10, 5))
    fmt.Println("Div(10,5):", mathlib.Div(10, 5))
    fmt.Println("Square(7):", mathlib.Square(7))
    fmt.Println("SquareRoot(49):", mathlib.SquareRoot(49))
     fmt.Println("Add(10,5):", mathlib.Add(10, 5))
    fmt.Println("Sub(10,5):", mathlib.Sub(10, 5))
    fmt.Println("Mul(10,5):", mathlib.Mul(10, 5))
    fmt.Println("Div(10,5):", mathlib.Div(10, 5))
    fmt.Println("Square(7):", mathlib.Square(7))
    fmt.Println("SquareRoot(49):", mathlib.SquareRoot(49))
     fmt.Println("Add(10,5):", mathlib.Add(10, 5))
    fmt.Println("Sub(10,5):", mathlib.Sub(10, 5))
    fmt.Println("Mul(10,5):", mathlib.Mul(10, 5))
    fmt.Println("Div(10,5):", mathlib.Div(10, 5))
    fmt.Println("Square(7):", mathlib.Square(7))
    fmt.Println("SquareRoot(49):", mathlib.SquareRoot(49))
     fmt.Println("Add(10,5):", mathlib.Add(10, 5))
    fmt.Println("Sub(10,5):", mathlib.Sub(10, 5))
    fmt.Println("Mul(10,5):", mathlib.Mul(10, 5))
    fmt.Println("Div(10,5):", mathlib.Div(10, 5))
    fmt.Println("Square(7):", mathlib.Square(7))
    fmt.Println("SquareRoot(49):", mathlib.SquareRoot(49))
     fmt.Println("Add(10,5):", mathlib.Add(10, 5))
    fmt.Println("Sub(10,5):", mathlib.Sub(10, 5))
    fmt.Println("Mul(10,5):", mathlib.Mul(10, 5))
    fmt.Println("Div(10,5):", mathlib.Div(10, 5))
    fmt.Println("Square(7):", mathlib.Square(7))
    fmt.Println("SquareRoot(49):", mathlib.SquareRoot(49))
    fmt.Println("OrderedList:", mathlib.OrderedList([]int{9, 3, 7, 1, 5}))

    fmt.Println("\n--- Heavy Math (may take time) ---")

    fmt.Println("Counting primes up to 50,000 ...")
    primes := mathlib.HeavyPrimeCount(50000)
    fmt.Println("Prime count:", primes)

    fmt.Println("\nComputing factorial of 1000 ...")
    fact := mathlib.BigFactorial(1000)
    fmt.Println("Factorial digit length:", len(fact.String()))

    fmt.Println("\nRunning trig series ...")
    trig := mathlib.TrigSeries(10_000_000)
    fmt.Println("Trig result:", trig)

    fmt.Println("\n===== APP COMPLETED SUCCESSFULLY =====")
}
