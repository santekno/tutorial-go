package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(str string) bool {
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	if reversedStr == str {
		fmt.Println("true")
		return true
	}
	fmt.Println("false")
	return false
}

func main() {
	fmt.Println("Masukkan kata/kalimat = ")
	var str string
	fmt.Scanln(&str)
	strToLower := strings.ToLower(str)
	IsPalindrome(strToLower)
}
