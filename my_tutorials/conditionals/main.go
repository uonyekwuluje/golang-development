package main

import (
  "fmt"
)

func main() {
  var cityName string
  cityName = "Bradford"
  fmt.Println("Welcome to => ",cityName)

  if cityName == "Bradford" {
    fmt.Println("You are in the right City")
  } else {
    fmt.Println("You are in the wrong City")
  }
}

