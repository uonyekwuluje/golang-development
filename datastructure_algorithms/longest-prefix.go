package main

import (
  "fmt"
)

func main() {
  //strArray := []string{"flower","flow","flight"}
  strArray := []string{"dog","racecar","car"}
  fmt.Println("Longest Preffix =>",longestCommonPrefix(strArray))
}

// find longest common prefix
func longestCommonPrefix(strs []string) string {
  var found string

  if len(strs) < 1 || len(strs) >= 200 {
    return ""
  }

  for _, item := range strs {
    if len(item) <= 0 || len(item) >= 200{
       return ""
    }
  }

  smallest := smallestLen(strs)

  for x:=1; x<=smallest; x++{
    z := 0
    for y:=0; y<len(strs); y++ {
       if strs[z][:x] != strs[y][:x] {
         return found
       }
    }
   found = strs[z][:x]
   z++
  }
  return found
}

// find Shortest String in list
func smallestLen(strArray []string) int {
  arrayLen := make([]int,len(strArray))
  for x:=0; x<len(strArray); x++{
    arrayLen[x] = len(strArray[x])
  }
  smallest := arrayLen[0]
  for x:=0; x<len(arrayLen); x++{
     if arrayLen[x] < smallest {
      smallest = arrayLen[x]
     }
  }
  return smallest
}
