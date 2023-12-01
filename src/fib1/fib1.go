package main

import "fmt"

func fib1(n int) int {
  if n < 2 {
    return n
  }

  return fib1(n-2) + fib1(n-1)
}

func main() {
  result := fib1(6)
  fmt.Printf("%d\n", result)
}
