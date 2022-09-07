package main

import (
	"fmt"
)
func Reverse(str string) (result string) {
   for _, v := range str {
      result = string(v) + result
   }
   return
}
func isPalindrome(input string) bool {
   if input != Reverse(input) {
      return false
   }
   return true
}
func main() {
   var input string
   fmt.Print("input to check Palindrome: ")
   fmt.Scanln(&input)
   result := isPalindrome(input)
	 if result == true {
      fmt.Println(true)
   } else {
      fmt.Println(false)
	 }

}