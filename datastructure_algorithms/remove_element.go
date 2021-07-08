package main

import (
  "fmt"
)

func main() {
  numArray := []int{0,1,2,2,3,0,4,2}
  removeElement(numArray,2)
}


//func removeElement(nums []int, val int) int {
func removeElement(nums []int, val int) {
  fmt.Println(nums)
  for x:=0; x<len(nums); x++ {
    for y:=0; y<len(nums); y++ {
         if nums[y] == val {
           nums[y] = nums[y+1]
         }
      }
    }
   fmt.Println(nums)
}
