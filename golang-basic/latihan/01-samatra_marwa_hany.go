package main

import (
	"fmt"
	"strings"
)

func main() {

	var inputString string
	fmt.Print("Masukan kata: ")
	fmt.Scanln(&inputString)
	var reverseString string = ""
	var length = len(inputString)

	for i := length - 1; i >= 0; i-- {
		reverseString = reverseString + string(inputString[i])
	}

	// Case in-sensitive comparision
	if strings.ToLower(inputString) == strings.ToLower(reverseString) {
		fmt.Println("The given string is Palindrome")
	} else {
		fmt.Println("The given string is NOT a Palindrome")
	}
}
