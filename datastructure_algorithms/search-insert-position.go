package main

import (
  "fmt"
)

func main() {
  numArray := []int{1,3}
  num := 3
  //numArray := []int{1,3,5,6}
  //num := 0
  fmt.Println(searchInsert(numArray,num))
}



func searchInsert(nums []int, target int) int {
   var found int
   for x:=0; x<len(nums); x++ {
     if (target == nums[x]) {
         found = x
         break
     } else if (target < nums[x]) {
         found = x
         break
     } else if (target > nums[len(nums)-1]) {
         found = len(nums)
         break
     }
   }
  return found
}
