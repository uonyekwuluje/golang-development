package main

import (
  "fmt"
  "os"
  "strconv"
  "./src"
)

var employee = map[string]int{"Mark":10,"Sunday":20}

func main() {
  for _, arg := range os.Args[1:] {
      t, err := strconv.ParseFloat(arg, 64)
      if err != nil {
         fmt.Fprintf(os.Stderr, "cf: %v\n", err)
         os.Exit(1)
      }
      f := tempconv.Fahrenheit(t)
      c := tempconv.Celsius(t)
      fmt.Printf("%s = %s, %s = %s\n",f, tempconv.FToC(f), c, tempconv.CToF(c))
  }

  name2 := []string {"uche","chinwe"}
  fmt.Println(name2)
  name2 = append(name2,"Emekea")
  fmt.Println(name2)
  fmt.Println(employee)
  fmt.Println(employee["Sunday"])
  farray()
}


func farray() {
   var nm [5]int
   fmt.Println("Testing Array")
   nm[0] = 4
   x := []int{3,4,5,6,7,9,23}
   fmt.Println(x)
   fmt.Println(x[2:5])
   x = append(x,89,90)
   fmt.Println(x)
}
