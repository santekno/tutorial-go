package main

import (
	"fmt"
	"strings"
)

func main() {

	var originalString string = "Madam"
	var reverseString string = ""
	var length = len(originalString)

	for i := length - 1; i >= 0; i-- {
		reverseString = reverseString + string(originalString[i])
	}

	// Case in-sensitive comparision
	if strings.ToLower(originalString) == strings.ToLower(reverseString) {
		fmt.Println("The given string is Palindrome")
	} else {
		fmt.Println("The given string is NOT a Palindrome")
	}
}
