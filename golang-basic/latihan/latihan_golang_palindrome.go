package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("input the wording: ")

	var palindromeInput string

	fmt.Scan(&palindromeInput)

	fmt.Println("inputted word is " + palindromeInput)
	var isPalindrome bool = isPalindromeWord(palindromeInput)

	fmt.Println(isPalindrome)
}

func isPalindromeWord(name string) bool {
	var chars []rune = []rune(name)
	var limitChecker int = int(len(chars) / 2)
	var charsCount int = len(chars)
	for i := 0; i < limitChecker; i++ {
		var leftChar string = string(chars[i])
		leftChar = strings.ToLower(leftChar)
		var rightChar string = string(chars[charsCount-(i+1)])
		rightChar = strings.ToLower(rightChar)
		// println(char)
		if leftChar != rightChar {
			return false
		}
	}
	return true
}
