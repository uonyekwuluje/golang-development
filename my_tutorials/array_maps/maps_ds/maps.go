package maps

import (
  "fmt"
)

func MapIntro() {
  fmt.Println("\nMap Intro")
  fmt.Println("------------")
  var fruit = map[string]int{}
  fmt.Println(fruit)
  fruit["Orange"] = 3
  fmt.Println(fruit)
  fruit["Mango"] = 4
  fruit["Guava"] = 6
  fruit["Pawpaw"] = 14
  fmt.Println(fruit)
  fmt.Println(fruit["Mango"])
}
