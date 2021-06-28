package main

import (
  "fmt"
)

func main() {
  num := []int{11,10,4,60,3,120,18,2,56,40}
  fmt.Println("Array => ",num)
  fmt.Println("Smallest => ",findSmallest(num))
  fmt.Println("Largest => ",findLargest(num))
}

// Find smallest number
func findSmallest(sm []int) int {
  small := sm[0]
  for x:=0; x<len(sm); x++{
   if sm[x] < small {
     small = sm[x]
   }
  }
  return small
}

// Find largest number
func findLargest(lg []int) int {
  large := lg[0]
  for x:=0; x<len(lg); x++{
   if lg[x] > large {
     large = lg[x]
   }
  }
  return large
}
