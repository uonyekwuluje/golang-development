package main

import (
  "fmt"
  "os"
  "strconv"
  "reflect"
)

func main() {
   decision1()
   decision2()
   loop1()
}


func decision1(){
  n1, _ := strconv.Atoi(os.Args[1])
  n2, _ := strconv.Atoi(os.Args[2])
  if (n1 > n2) {
     fmt.Printf("%d is greater than %d\n",n1,n2)
  } else {
    fmt.Printf("%d is greater than %d\n",n2,n1) 
  }
}


func decision2() {
  n1, _ := strconv.Atoi(os.Args[1])
  fmt.Println("Decision 2")
  fmt.Println(".............")
  switch n1{
  case 2,3,4:
    fmt.Printf("Less than 5\n")
  case 5:
    fmt.Printf("Just 5\n")
  case 6,7,8,9:
    fmt.Printf("Greater than 5\n")
  default:
    fmt.Printf("Something Else\n")
  }
}


func loop1(){
  k := 1
  fmt.Println("Testing Loops")
  fmt.Println(".............")
  for; k <= 10; k++ {
    fmt.Println(k)
  }

  var n int = 0
  for range "Hello"{
   fmt.Printf("%d and H\n",n)
   n++
  }
  fmt.Println(reflect.ValueOf(n).Kind())
}
