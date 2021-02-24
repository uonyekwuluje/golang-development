package arrays

import ("fmt")


func StringProcess() {
  var names[3] string
  var string_test = "Welcome Fizz Buzz"

  names[0] = "ball"
  fmt.Print("This is :> ",names[1])

  fmt.Println(string_test)
  fmt.Println(string_test[:3])
  fmt.Println(string_test[5:7])
  fmt.Println(len(string_test))
}


func IntArrayBasics() {
  fmt.Println("\nArray Basics")
  fmt.Println("---------------")
  var num = []int{3,4}
  fmt.Printf("Current Array is of Lenght => %d\n",len(num))
  num = append(num, 56)
  fmt.Printf("Current Array is of Lenght => %d\n",len(num))
  a := []int{45,23}
  b := []int{145,423}
  num = append(num, a...)
  num = append(num, b...)
  fmt.Printf("Current Array is of Lenght => %d\n",len(num))
}

func StringToArray() {
  fmt.Println("\nArray Manipulation")
  fmt.Println("--------------------")

  var tstring = "Hello Golang"
  var tarry = []string{}
  fmt.Printf("New String %s is of Lenght %d\n",tstring,len(tstring))
  for x:=0; x<len(tstring); x++ {
    fmt.Println("Char => ",string(tstring[x]))
    tarry = append(tarry, string(tstring[x]))
  }
  fmt.Println("New Array => ",tarry)
  fmt.Println(tarry[0])
}
