package main

import (
  "fmt"
  "math/big"
  "time"
)

func main() {
  fmt.Println("🚀 Starting 5-minute real-time compute...")
  start := time.Now()
  limit := 305 * time.Second
  a, b := big.NewInt(0), big.NewInt(1)

  for time.Since(start) < limit {
    a.Add(a, b)
    a, b = b, a
    if int(time.Since(start).Seconds())%30 == 0 && a.BitLen()%100 == 0 {
      fmt.Printf("⏱️ Running... %v elapsed.\n", time.Since(start).Round(time.Second))
    }
  }
  fmt.Printf("✅ Success: Finished after %v\n", time.Since(start))
}
