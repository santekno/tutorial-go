package main

import (
	"fmt"
	"strings"
)

func palindrome(input string) string {
	var reverseString string = ""
	var length = len(input)

	for i := length - 1; i >= 0; i-- {
		reverseString = reverseString + string(input[i])
	}

	// Case in-sensitive comparision
	if strings.ToLower(input) == strings.ToLower(reverseString) {
		return "palindrome"
	} else {
		return "not palindrome"
	}
}

func main() {

	var inputString string
	fmt.Print("Masukan kata: ")
	fmt.Scanln(&inputString)
	palindrome(inputString)

}
