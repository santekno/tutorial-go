package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Input Word: ")
	fmt.Scanln(&input)
	fmt.Println(isVowelOrConsonant(input))
}

func isVowelOrConsonant(val string) string {
	switch strings.ToLower(val) {
	case "a", "i", "u", "e", "o":
		return "Vokal"
	default:
		return "Konsonan"
	}
}