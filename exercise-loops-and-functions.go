package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) (float64, int) {
  z := float64(1)
  pz := float64(0)
  i := 0

  for ; math.Abs(pz-z) > 0.0001; i++ {
    pz = z
    z = z - ((z*z - x) / 2*z)
  }

  return z, i
}

func main() {
    z, i := Sqrt(2)
    fmt.Println("loop:", i)
    fmt.Println(z)
    fmt.Println(math.Sqrt(2) - z)
}
