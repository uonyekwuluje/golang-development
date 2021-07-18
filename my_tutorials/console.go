package main

import (
  "fmt"
)

func main() {
  vint := []int{5,7,8,9}
  for _,x := range vint {
      fmt.Println(x)
  }
}
