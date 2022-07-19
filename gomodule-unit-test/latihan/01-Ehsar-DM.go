package main

import (
	"fmt"
)

func main() {
	// Soal 1
	status := false
	fmt.Print("Enter a word : ")
	var word string
	fmt.Scanf("%s", &word)

	reverse := func(word string) (result string) {
		for _, v := range word {
			result = string(v) + result
		}
		return
	}

	if word == reverse(word) {
		status = true
	}

	fmt.Println(status)
}

func polindrome(word string) bool {
	status := false

	reverse := func(word string) (result string) {
		for _, v := range word {
			result = string(v) + result
		}
		return
	}

	if word == reverse(word) {
		status = true
	}

	return status
}
