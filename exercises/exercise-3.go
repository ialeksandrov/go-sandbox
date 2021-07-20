package main

import (
 "golang.org/x/tour/wc"
 "strings"
)

func WordCount(s string) map[string]int {
  result := make(map[string]int)
  words := strings.Fields(s)
  for _, word := range words {
    result[word] += 1
  }
  return result
}
func main() {
  wc.Test(WordCount)
}
