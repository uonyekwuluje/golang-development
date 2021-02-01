package main

import (
   "fmt"
   "reflect"
   "io/ioutil"
   "log"
)

// Global Variables
var (
    product  = "Mobile"
    quantity = 50
    price    = 50.50
    inStock  = true
)

// Global Constant
const (
    SUMINIT = 5
)

func main(){
   fmt.Printf("Welcome To Go Basics\n")
   fmt.Printf("----------------------\n")
   fmt.Println("Const is => ",SUMINIT)
   dataTypes()
   globalTest()
   operations()
   sysops1()
}


func sysops1() {
   fmt.Printf("\nFile Operations\n")
   fmt.Printf("----------------------\n")
   files, err := ioutil.ReadDir("./")
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
            fmt.Println(f.Name())
    }
}


func operations() {
  fmt.Printf("\nTesting Operations\n")
  fmt.Printf("----------------------\n")
  var x, y, z = 7,5,8
  fmt.Println("You Just Entered ",x,y)
  fmt.Println("7 / 5 =>",x/y," remainder ",x%y)
  fmt.Println(x < y && x > z)
  fmt.Println(x < y || x > z)
  fmt.Println(!(x == y && x > z))
}


func dataTypes() {
   fmt.Printf("\nTesting Data Types\n")
   fmt.Printf("------------------\n")
    var i = 10
    var s = "Canada"
    name := "John Doe"
    var n1 int = 45
    var n2 float64 = 45.456

    fmt.Println(reflect.TypeOf(name))
    fmt.Println(reflect.TypeOf(i))
    fmt.Println(reflect.TypeOf(s))
    fmt.Println(reflect.TypeOf(n2))
    fmt.Println("Testing n1 ",n1)
}


func globalTest() {
    fmt.Printf("\nTesting Global Variables\n")
    fmt.Printf("------------------------\n")
    fmt.Println(quantity)
    fmt.Println(price)
    fmt.Println(product)
    fmt.Println(inStock)
}
