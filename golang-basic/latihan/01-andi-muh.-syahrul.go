package main

import (
	"fmt"
	"strings"
)

func isPalindrome(input string) bool {

	value := strings.ToLower(input)

	for i := 0; i < len(value)/2; i++ {
		if value[i] != value[len(value)-i-1] {
			return false
		}
	}

	return true

}

func main() {

	var str string

	fmt.Printf("Type a word: ")
	fmt.Scanf("%s\n", &str)

	result := isPalindrome(str)

	fmt.Printf("%v", result)

}
