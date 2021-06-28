package main

import (
  "fmt"
)

func main() {
   xs := []int{4,6,8,2,8}
   total := 0
   for _, v := range xs {
      total += v
   }
  fmt.Println("Total => ",total)
}
