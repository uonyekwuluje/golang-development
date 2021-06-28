package main

import (
  "fmt"
)

func main() {
  x := 45
  fmt.Println(testParm(x))
}


func testParm(y int) int {
  return y+1
}
