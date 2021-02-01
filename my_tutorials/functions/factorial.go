package main

import "fmt"

func main() {
  var num int = 5
  fmt.Printf("The Factorial of %d => %d",num,factorial(num))
}

func factorial(x int) int {
  if x == 0 {
    return 1
  }
  return x * factorial(x-1)
}
