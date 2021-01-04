package main

import (
  "math"
  "fmt"
)


func add(x int, y int) int {
  return x + y
}

func concat(x string, y string) string {
  return x + y
}

func multiply(x int, y int) int {
  return x * y
}

func power(x float64, y float64) float64 {
  return math.Pow(x, y)
}

func main() {
  fmt.Println(add(42, 13))
  fmt.Println(concat("Hello", " World!"))
  fmt.Println(multiply(2, 3))
  fmt.Println(power(-4, 2))
}
