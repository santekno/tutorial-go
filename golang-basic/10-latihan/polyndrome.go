package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Scanf("%s", &name)

	fmt.Println(isPalindrome(name))
}

func isPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
