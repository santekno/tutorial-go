package main

import (
	"fmt"
)

func main() {
	var str string

	getInput(&str)

	if checkPalindrome(str) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

func getInput(str *string) {
	fmt.Println("\nCheck if a string is palindrome")
	fmt.Print("===============================", "\n\n")
	fmt.Print("Input some string: ")
	fmt.Scan(str)
}

func checkPalindrome(str string) (palindrome bool) {
	var left int = 0
	var right int = len(str) - 1

	for left < right {
		if str[left] != str[right] {
			return
		}

		left++
		right--
	}

	return true
}
