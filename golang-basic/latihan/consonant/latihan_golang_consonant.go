package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("input the wording: ")

	var consonantInput string

	fmt.Scan(&consonantInput)

	fmt.Println("inputted word is " + consonantInput)
	var isConsonant bool = isConsonant(consonantInput)

	fmt.Println(isConsonant)
}

func isConsonant(character string) bool {
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	if isAlpha(character) {
		var charToCheck = strings.ToLower(character)
		var consonantCharacters []string = []string{"a", "i", "u", "e", "o"}

		for _, element := range consonantCharacters {
			if charToCheck == element {
				return false
			}
		}

		return true
	} else {
		return false
	}
}
