package main

import (
    "fmt"
)

func isPalindrome(n int) {
    orig := n
    new_int := 0
    for n > 0 {
        remainder := n % 10
        new_int *= 10
        new_int += remainder
        n /= 10
    }

    if orig == new_int {
      fmt.Println("n => ",orig," new_int => ",new_int)
    } else if orig != new_int {
      fmt.Println("n => ",orig," new_int => ",new_int)
    }
}

/*func reverse_int(n int) int {
    new_int := 0
    for n > 0 {
        remainder := n % 10
        new_int *= 10
        new_int += remainder
        n /= 10
    }
    return new_int
}*/

func main() {
    isPalindrome(123456)
    isPalindrome(100)
    isPalindrome(1001)
    isPalindrome(131415)
    isPalindrome(1357)
    isPalindrome(-121)
    isPalindrome(10)
}
