package main

import (
  "code.google.com/p/go-tour/wc"
  "strings"
)

func WordCount(s string) map[string]int {
  counter := make(map[string]int)
  for _, w := range strings.Fields(s) {
    counter[w] += 1
  }
  return counter
}

func main() {
  wc.Test(WordCount)
}
