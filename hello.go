package main

import (
  "fmt"
)

func main() {
   var n1,n2 string
   n1 = "Hello"
   n2 = "World"
   fmt.Println("String Lenght for ",n1," and ",n2," => ",len(n1))
   fmt.Println(len(n1) == len(n2))
}
