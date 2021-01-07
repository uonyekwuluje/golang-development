package main

import ("fmt")

func main() {
  var names[3] string
  var string_test = "Welcome Fizz Buzz"

  names[0] = "ball"
  fmt.Println("This is :> ",names[1])

  fmt.Println(string_test)
  fmt.Println(string_test[:3])
  fmt.Println(string_test[5:7])
  fmt.Println(len(string_test))
}
