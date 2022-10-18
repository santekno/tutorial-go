package main

import (
	"fmt"
	"strings"
)

func main() {
	// check the input is palindrome or not
	inputString := "kapak"
	v := isPalindrome(inputString)
	fmt.Println(v)

	// chceck konsonan
	var inputKonsonan string
	fmt.Println("Input the character: ")
	fmt.Scanln(&inputKonsonan)
	k := isKonsonan(inputKonsonan)
	fmt.Println(k)
}

func isPalindrome(inputString string) bool {
	splittedString := strings.Split(inputString, "")
	for j := len(splittedString) / 2; j > 0; j-- {
		if splittedString[len(splittedString)-j] != splittedString[j-1] {
			return false
		}
	}
	return true
}

func isKonsonan(inputString string) bool {
	var konsonan = map[string]bool{
		"a": true,
		"i": true,
		"u": true,
		"e": true,
		"o": true,
	}
	inputString = strings.ToLower(inputString)
	if len(inputString) > 1 {
		panic("only one character!")
	}
	_, isExist := konsonan[inputString]
	return isExist
}
