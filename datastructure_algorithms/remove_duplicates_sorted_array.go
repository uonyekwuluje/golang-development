package main

import (
  "fmt"
)

func main() {
  nm1 := [][]int{
    {6,8,3,4,5},
    {4,9,0,7,23},
  }
  fmt.Println("This is good ",len(nm1))

  //nm2 := []int{1,1,2}
  nm2 := []int{0,0,1,1,1,2,2,3,3,4}
  removeDuplicates2(nm2)
  fmt.Println(removeDuplicates(nm2))
}


func removeDuplicates2(nums []int) {
  var x,y,z int

  fmt.Println("This is good ",nums)
  for x=0; x<len(nums); x++{
    for y=1; y<len(nums); y++ {
       if nums[y] == nums[y-1] {
         for z=y; z<len(nums); z++ {
            nums[z-1] = nums[z]
        }
        nums[z-1] = 0
       }
    }
  }
  fmt.Println("This is good ",nums)
}


func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    j := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[i-1] {
            nums[j] = nums[i]
            j++
        }
    }
    fmt.Println(nums[:j]) //nums[:j] is the new deduplicated array
    return j
}
