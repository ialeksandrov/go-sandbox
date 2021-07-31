package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
  total, nextTotal := 0, 1
  return func() int {
    result := total
    total, nextTotal = nextTotal, nextTotal + result
    return result
  }
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
