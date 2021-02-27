package main

import (
  "fmt"
  "strconv"
)

func main() {
  var num int = 679236783
  var numArray []int64
  iLen := intLen(num)
  fmt.Printf("Integer [%d] of Lenght [%d] of Type [%T]\n",num,iLen,num)

  for x:=1; x<=iLen; x++ {
     n := strconv.Itoa(num)[x-1:x]
     nm, _ := strconv.ParseInt(n, 10, 64)
     numArray = append(numArray, nm)
  }
  fmt.Printf("New Array %d of Type %T\n",numArray,numArray)
}

/* return lenght of integer */
func intLen(x int) int {
  var count int = 0
  for x != 0 {
    x /= 10
    count += 1
  }
  return count
}
