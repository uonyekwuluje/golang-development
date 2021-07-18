package main

import (
  "fmt"
)

func main() {
  nd := []int{1,2,3}
  plusOne(nd)
}


func plusOne(digits []int) []int {
   fmt.Println("Original =>",digits)
   fmt.Println(digits[len(digits)-1])
   digits[len(digits)-1] = digits[(len(digits)-1)] + 1
   fmt.Println("Original Plus 1 =>",digits)
}
