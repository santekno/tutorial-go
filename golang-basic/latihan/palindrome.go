package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("LEVEL"))
	fmt.Println(isPalindrome("TANAH"))
	fmt.Println(isPalindrome("KATAK"))
	fmt.Println(isPalindrome("KASUR"))
	fmt.Println(isPalindrome("PALINDROME"))
	fmt.Println(isPalindrome("MADAM"))
}

func isPalindrome(str string) bool {
	word := len(str)

	for i := range str {
		if strings.ToLower(string(str[i])) != strings.ToLower(string(str[word-i-1])) {
			return false
		}
	}
	return true
}