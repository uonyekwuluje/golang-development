package main

import (
  "fmt"
)

type Person struct {
  FirstName   string
  LastName    string
  age         int
  Address
}

type Address struct {
  Street      string
  HouseNumber int
  ZipCode     int
  State       string
}


func main() {
  p := Person{"Uche","Onyekwuluje",45,Address{"Davis Street",4,34567,"MA"}}
  fmt.Println(p.FirstName)
  fmt.Println("Zip Code ",p.ZipCode)
}
