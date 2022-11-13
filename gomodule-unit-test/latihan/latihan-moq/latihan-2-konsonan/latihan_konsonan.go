package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Scan(&input)

	result := (consonant(input))
	if result {
		fmt.Println("Konsonan")
	} else {
		fmt.Println("Vokal")
	}
}

func consonant(inp string) bool {
	var out string
	out = strings.ToLower(inp)

	var vocal_char []string = []string{"a", "i", "u", "e", "o"}

	for _, element := range vocal_char {
		if out == element {
			return false
		}
	}

	return true
}
