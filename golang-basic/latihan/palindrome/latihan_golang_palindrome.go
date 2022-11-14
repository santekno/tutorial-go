package main

import (
	"fmt"
	"strings"

	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func main() {
	fmt.Println("input the wording: ")

	var palindromeInput string

	fmt.Scan(&palindromeInput)

	fmt.Println("inputted word is " + palindromeInput)

	fmt.Println(quote.Hello())
	fmt.Println(quoteV3.HelloV3())

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
