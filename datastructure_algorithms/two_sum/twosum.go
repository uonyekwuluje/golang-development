package main

import (
  "fmt"
)

func main() {
  numset1 := []int{2,5,5,11}
  fmt.Println(twoSum(numset1, 10))
}

func twoSum(nums []int, target int) []int {
  returnArray := []int{}
  arrayLen := len(nums)
  for x:=0; x < arrayLen - 1; x++ {
    for y:=x+1; y < arrayLen; y++ {
      //fmt.Printf("%d + %d = %d\n",nums[x],nums[y],(nums[x]+nums[y]))
      if nums[x] + nums[y] == target {
        returnArray = append(returnArray,x)
        returnArray = append(returnArray,y)
        return returnArray
      }
    }
  }
  return returnArray
}
