package main

import (
	"fmt"
	"strings"
)

func IsPalindrom(str string) bool {
	str = strings.ToLower(str)
	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var word string
	fmt.Println("Masukan Kata:")
	fmt.Scan(&word)
	palindrom := IsPalindrom(word)
	if palindrom == true {
		fmt.Println("is Palindrome")
	} else {
		fmt.Println("is not Palindrome")
	}
}
