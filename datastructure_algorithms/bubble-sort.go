package main

import (
  "fmt"
)

func main() {
   x := []int{
      48,96,86,9,
      57,17,63,70,
      37,34,83,27,
      68,19,97,82,
   }
   fmt.Println(x)

   for curr:=0; curr<len(x); curr++ {
     for inner:=1; inner<len(x); inner++ {
         if x[inner] < x[inner - 1] {
           tmp := x[inner]
           x[inner] = x[inner - 1]
           x[inner - 1] = tmp
         }
     }
   }
   fmt.Println(x)
}
