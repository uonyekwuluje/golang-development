package main

import (
   "fmt"
   "reflect"
   "strconv"
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
