// Convert Minutes to Seconds
package main

import (
  "fmt"
)

func main() {
   min := []int{5,3,2}
   for x:=0; x<len(min); x++ {
      fmt.Printf("convert(%d) -> %d\n",min[x],convert(min[x]))
   }
}


func convert(x int) int {
  return x * 60
}
