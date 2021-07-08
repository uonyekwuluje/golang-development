package main

import (
 "fmt"
)

func main() {
  alphaCount := make(map[string]int)
  fmt.Println("Testing Map")
  fmt.Println("-------------")
  alphaCount["The"] = 4
  alphaCount["Cost"] = 2
  alphaCount["Fox"] = 1
  for _,y := range alphaCount {
    fmt.Println(y)
  }
  delete(alphaCount,"Cost")
  fmt.Println(alphaCount)
}
