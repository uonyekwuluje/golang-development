package main

import (
   "fmt"
   "strings"
)

func main() {
    str := "Hello World"
    fmt.Println("Lenght of Last Word =>",lengthOfLastWord(str))
}

func lengthOfLastWord(s string) int {
    if len(s) == 0 || len(s) >= 10*10*10*10 {
       return 0
   }

    s1 := strings.Split(s," ")
    return len(s1[len(s1)-1])
}
