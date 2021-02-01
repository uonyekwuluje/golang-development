/* 
Testing out command line arguments
./console-code.go <arg1> <arg2>
*/
package main

import (
  "fmt"
  "os"
)

func main() {
   fmt.Println("Testing Console Code")
   fmt.Println("First Argument =>",os.Args[1])
   fmt.Println("Second Argument =>",os.Args[2])
   fmt.Println("All Arguments =>",os.Args[:])
}
