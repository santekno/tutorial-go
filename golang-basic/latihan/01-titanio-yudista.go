package main

import (
	"fmt"
)

func main() {
	var word string
	//! Input word via terminal
	fmt.Scanf("%s", &word)
	fmt.Println(isPalindrome(word))
}

func isPalindrome(substr string) bool {
	result := []byte{}
	//! looping decrement
	for i := len(substr) - 1; i >= 0; i-- {
		//! Append to slice result
		result = append(result, substr[i])
	}

	//! compare result and parameter input
	return substr == string(result)
}
