package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("madam"))  // true
	fmt.Println(isPalindrome("achmad")) // false
	fmt.Println(isPalindrome("m"))      // true
	fmt.Println(isPalindrome("MaDam"))  // true
	fmt.Println(isPalindrome("-+_+-"))  // true
}

func isPalindrome(word string) bool {
	for i := 0; i < len(word); i++ {
		j := len(word) - 1 - i
		if strings.ToLower(string(word[i])) != strings.ToLower(string(word[j])) {
			return false
		}
	}
	return true
}
