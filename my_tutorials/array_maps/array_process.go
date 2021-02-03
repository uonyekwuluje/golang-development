package main

import ("fmt")

func main() {
  stringProcess()
  dictMaps()
  findSmallest()
  dictMaps1()
}

func stringProcess() {
  var names[3] string
  var string_test = "Welcome Fizz Buzz"

  names[0] = "ball"
  fmt.Print("This is :> ",names[1])

  fmt.Println(string_test)
  fmt.Println(string_test[:3])
  fmt.Println(string_test[5:7])
  fmt.Println(len(string_test))
}

func dictMaps() {
  numfreq := make(map[string]int)
  numfreq["Nigeria"] = 80
  numfreq["USA"] = 5
  numfreq["Zimbabwe"] = 5
  fmt.Println(numfreq)
  delete(numfreq,"USA")
  fmt.Println(numfreq)
  el,ok := numfreq["USA"]
  fmt.Println(el,ok)
}

func findSmallest() {
  nums := []int{4,1,6,9,90}
  smallest := nums[0]
  for _, element := range nums{
    if element < smallest{smallest = element}
  }
  fmt.Println("Smallest Number => ",smallest)
}



func dictMaps1() {
  strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
  for index, element := range strDict {
     fmt.Println("Index :", index, " Element :", element)
  }

  for key := range strDict {
    fmt.Println(key)
  }

  for _, value := range strDict {
    fmt.Println(value)
  }
}
