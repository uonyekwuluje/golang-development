package main

import (
  "fmt"
  "math"
  "unicode"
)

func main() {
   str := "a"
   fi := ""
   fmt.Println(strStr(str,fi))
}

func strStr(haystack string, needle string) int {
  if len(haystack) < 0 || len(haystack) > (5 * int(math.Pow(10,4))) {
    return 0
  }

  for _,ch := range haystack {
   if (unicode.IsUpper(ch) || !unicode.IsLetter(ch)) {
     return 0
   }
  }

  if len(haystack) == 0 && len(needle) == 0 {
    return 0
  }

   for x:=0; x<len(haystack); x++ {
     if (x+2) > len(haystack) {
       break
     } else {
       //fmt.Println(haystack[x:x+2])
       if haystack[x:x+2] == needle {
        //fmt.Println("Found At Index =>",x)
        return x
       }
    }
   }
  return -1
}

