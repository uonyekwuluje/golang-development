package main

import (
  "os"
  "fmt"
)

func main() {
  cwd,_ := os.Getwd()
  fmt.Println(cwd)
}
