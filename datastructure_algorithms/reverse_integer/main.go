package main

import (
  "fmt"
)

func main() {
  var num int = -6792367830
  fmt.Println("Integer =>",num)
  fmt.Println("Reverse =>",reverse(num))
}


/* Reverse Integer */
func reverse(x int) int{
  var MaxInt int32 = 2147483647
  if x > int(MaxInt) || x < - int(MaxInt) { return 0 }

  new_int := 0
  for x > 0 {
    remainder := x % 10
    new_int *= 10
    new_int += remainder
    x /= 10
  }

  if new_int > int(MaxInt) { return 0 }
  if (x < 0) {
        return -new_int
  }

  return new_int
}
