package main

import (
  "fmt"
)

func main() {
  var num int = -6792367830
  fmt.Println("Integer =>",num)
  fmt.Println("Reverse1 =>",reverse1(num))
  fmt.Println("Reverse2 =>",reverse2(num))
}


/* Reverse Integer */
func reverse1(x int) int{
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

func reverse2(x int) int {
	k := 1
	limit := 2147483648
	if x < 0 {
		k = -1
		x *= k
	}

	reverse := 0
	for x > 0 {
		reverse = (reverse * 10) +  x % 10
		x /= 10
	}
	if reverse > limit || reverse < -limit {
		return 0
	}
	return reverse * k
}
