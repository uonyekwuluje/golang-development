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
   fmt.Printf("\n\n")

   intNums := []int{34,56,78,90,25,67}
   testIntFunc(intNums)
   fmt.Printf("\n\n")

   fmt.Println("Sum => ",testMultIntFunc(34,56,78,90,25,67))
}


func func_dec1(n1, n2 int) int {
  fmt.Println("This function returns the biggest number")
  if (n1 > n2) {
    return n1
  } else {
    return 0
  }
}

func swap(x, y string) (string, string) {
  fmt.Println("This function Sawps arguments")
   return y, x
}

func testIntFunc(nums []int) {
  fmt.Println("Testing Function with Array Argument")
  fmt.Println(nums)
}

func testMultIntFunc(nums ...int) int {
  fmt.Println("Testing Function with Multiple Argument")
  total := 0
  for _,val := range nums {
    total = total + val
  }
  return total
}
