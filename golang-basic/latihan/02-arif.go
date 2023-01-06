package main

import (
	"fmt"
	"strings"
)

func main() {
	var kata string

	fmt.Scan(&kata)

	kata2 := kata[0]

	result := checkConsonant(string(kata2))

	fmt.Print(result)
}

func checkConsonant(ch string) string {
	ch = strings.ToLower(ch)

	if ch == "a" || ch == "i" || ch == "u" || ch == "e" || ch == "o" {
		return "Vokal"
	}

	return "Konsonan"
}

func checkConsonantWithPackage(input string) string {
	vowels := "aeiou"

	if strings.Contains(vowels, input) {
		return "Vokal"
	}

	return "Konsonan"
}
