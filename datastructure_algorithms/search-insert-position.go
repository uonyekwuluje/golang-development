package main

import (
  "fmt"
)

func main() {
  numArray := []int{1,3,5,6}
  num := 5
  searchInsert(numArray,num)
}



//func searchInsert(nums []int, target int) int {
func searchInsert(nums []int, target int) {
    fmt.Println("Numbers =>",nums)
    fmt.Println("Target =>",target)
}
