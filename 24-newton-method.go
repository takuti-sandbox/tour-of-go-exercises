package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) (z float64) {
  eps := 1e-3
  z = 1.0
  z_prev := z
  for {
    z = z - (z * z - x) / (2 * z)
    if math.Abs(z_prev - z) < eps { break }
    z_prev = z
  }
  return
}

func main() {
  fmt.Println("Sqrt(2) =", Sqrt(2))
  fmt.Println("math.Sqrt(2) =", math.Sqrt(2))
}
