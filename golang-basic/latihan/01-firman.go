package main

import (
	"fmt"
	"strings"
)

func isPalindrome(str *string) bool {
	*str = strings.ToLower(*str)
	for i := 0; i < len(*str); i++ {
		j := len(*str) - 1 - i
		if (*str)[i] != (*str)[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Print("Masukkan sebuah Kata: ")
	var input string
	fmt.Scanln(&input)
	fmt.Println(isPalindrome(&input))
}
