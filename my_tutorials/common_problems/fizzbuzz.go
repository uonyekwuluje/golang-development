package main

import (
  "fmt"
)

func main() {
  fmt.Println("FizzBuzz 1")
  fmt.Println("-----------")
  x:= 100
  for y:=1;y<=x;y++ {
   if (y%3 == 0) && (y%5 == 0) {
      fmt.Println(y," FizzBuzz")
   } else if (y%5 == 0) {
      fmt.Println(y," Buzz")
   } else if (y%3 == 0) {
      fmt.Println(y," Fizz")
   }
  }
  fizzbuzz2()
}

func fizzbuzz2() {
  fmt.Println("FizzBuzz 2")
  fmt.Println("-----------")

    for i := 1; i <= 100; i++ {

        if i%3 == 0 {
            // Multiple of 3
            fmt.Printf("fizz")
        }
        if i%5 == 0 {
            // Multiple of 5
            fmt.Printf("buzz")
        }

        if i%3 != 0 && i%5 != 0 {
            // Neither, so print the number itself
            fmt.Printf("%d", i)
        }

        // A trailing new line (so both fizz + buzz can be printed on the same line)
        fmt.Printf("\n")

    }
}
