package main

import (
  "fmt"
)

func main() {
   var num1 int = 400
   var num2 int = 14
   fmt.Printf("Max value is: %d\n", func_dec1(num1,num2))
   var st1 string = "Addition"
   var st2 string = "Subtraction"
   fmt.Printf("Before Swap: %s\t%s\n", st1,st2)
   st1, st2 = swap(st1, st2)
   fmt.Printf("Before Swap: %s\t%s\n", st1,st2)
}


func func_dec1(n1, n2 int) int {
  if (n1 > n2) {
    return n1
  } else {
    return 0
  } 
}

func swap(x, y string) (string, string) {
   return y, x
}
