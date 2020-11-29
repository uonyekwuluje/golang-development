// function to add and return sum
package main

import (
   "fmt" 
   "bufio" 
   "os"
)

func main() {
   inputCml1()
   inputCml2()
   operatorsInc()
}


func inputCml1() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter Text: [Hit Return to end] ")
  text, _ := reader.ReadString('\n')
  fmt.Println("String Entered =>",text)
}

func inputCml2() {
  var num1,num2 int
  fmt.Print("Enter Number: [Hit Return to end] ")
  fmt.Scanf("%d",&num1)
  fmt.Print("Enter Another Number: [Hit Return to end] ")
  fmt.Scanf("%d",&num2)
  fmt.Printf("The sum of %d and %d = %d",num1,num2,(num1+num2))
}

func operatorsInc(){
  const NUM1 int = 10
  const NUM2 int = 10
  fmt.Println(NUM1 == NUM2)
}
