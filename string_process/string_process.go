package main

import ("fmt"; "bufio"; "os")

func main() {
   stringProc1()
   stringProc2()
}


func stringProc1() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter Text: [Hit Return to end] ")
  text, _ := reader.ReadString('\n')
  fmt.Println("String Entered =>",text)
}

func stringProc2() {
  var greet = "String Processing"
  fmt.Println("String Test => ",greet)
  fmt.Println("String Test => ",greet," is of Lenght => ",len(greet))
}
