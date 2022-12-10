package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputString string
	fmt.Println("Check Palindrome : ")
	fmt.Scanf("%s", &inputString)
	fmt.Println(isPalindrome(inputString))
}

func isPalindrome(inputString string) string {
	inputString = strings.ToLower(inputString)
	for i := 0; i < len(inputString)/2; i++ {
		if inputString[i] != inputString[len(inputString)-1-i] {
			return "Not Palindrome"
		}
	}
	return "Palindrome"
}
