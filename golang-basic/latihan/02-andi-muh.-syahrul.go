package main

import (
	"fmt"
	"unicode"
)

func isVowel(input rune) string {

	var char rune = unicode.ToLower(input)

	if (char == 'a') || (char == 'e') || (char == 'i') ||
		(char == 'o') || (char == 'u') {

		return "Vokal"

	} else {

		return "Konsonan"

	}

}

func main() {

	var char rune

	fmt.Printf("Type a char: ")
	fmt.Scanf("%c\n", &char)

	result := isVowel(char)

	fmt.Printf("%s", result)

}
