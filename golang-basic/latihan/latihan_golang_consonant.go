package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("input the wording: ")

	var consonantInput string

	fmt.Scan(&consonantInput)

	fmt.Println("inputted word is " + consonantInput)
	var isConsonant bool = isConsonant(palindromeInput)

	fmt.Println(isConsonant)
}

func isConsonant(character string) bool {
	var charToCheck = strings.ToLower(character)
	var consonantCharacters []string = []string{"a", "i", "u", "e", "o"}

	for _, element := range consonantCharacters {
		if charToCheck == element {
			return true
		}
	}

	return false
}
