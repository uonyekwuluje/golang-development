package main

import (
  "fmt"
)

func main() {
   arr := []int{2,3,4,10,15,23,67}
   x := 15
   fmt.Println("The Array => ",arr,"Find Item",x)

   result := binarySearch(arr,0,len(arr)-1,x)
   if result != -1 {
     fmt.Println("Found Item In Position ",result)
   } else {
     fmt.Println("Item Not Found")
   }
}


func binarySearch(arr []int, left int, right int, item int) int {
  if right >= left {
    mid := left + (right - left) / 2
    if arr[mid] == item {
       return mid
    } else if arr[mid] > item {
       return binarySearch(arr,left,mid-1,item)
    } else {
       return binarySearch(arr,mid+1,right,item)
    }
  } else {
      return -1
  }
}
