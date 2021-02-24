package main

import "fmt"

func swap(x, y string) (string, string) {
  return y, x
}

func swap_int(x, y int) (int, int) {
  return y, x
}

func main() {
  a, b := swap("Hello ", "World")
  fmt.Println(a, b)

  a1, b1 := swap_int(23, 25)
  fmt.Println(a1, b1)
}
