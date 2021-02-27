package main

import (
  "fmt"
  "strconv"
)

func main() {
  var num int = 679236783
  var numArray []int64
  iLen := intLen(num)
  fmt.Println("Integer =>",num)
  fmt.Println("Interger Lenght =>",iLen)

  for x:=iLen; x>=1; x-- {
     //fmt.Print(strconv.Itoa(num)[x-1:x])
     n := strconv.Itoa(num)[x-1:x]
     nm, _ := strconv.ParseInt(n, 10, 64)
     fmt.Printf("%d of type %T\n",nm,nm)
     numArray = append(numArray, nm)
  }
  fmt.Println(numArray)
  fmt.Printf("%T\n",numArray)
  fmt.Println(reverse(num))

  /*
  for x:=1; x<=iLen; x++ {
     fmt.Println("Current Index",strconv.Itoa(num)[x-1:x])
  }

  fmt.Println("\nRevert Numbers")
  for x:=iLen; x>=1; x-- {
     fmt.Println("Current Index",strconv.Itoa(num)[x-1:x])
  }
  */
}


/* Reverse Integer */
func reverse(x int) int{
  new_int := 0
  for x > 0 {
    remainder := x % 10
    new_int *= 10
    new_int += remainder
    x /= 10
  }
  return new_int
}

/* Compute and return lenght of integer */
func intLen(x int) int {
  var count int = 0
  for x != 0 {
    x /= 10
    count += 1
  }
  return count
}
