package main

import (
	"fmt"
)

func isContains(listAlphabet []string, alphabet string) bool {
	for _, v := range listAlphabet {
		if v == alphabet {
			return true
		}
	}
	return false
}

func isConsonant(input string) bool {
	listConsonant := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}
	return isContains(listConsonant, input)
}

func isVowel(input string) bool {
	listVowel := []string{"a", "e", "i", "o", "u"}
	return isContains(listVowel, input)
}

func checkVowelOrConsonant(input string) string {
	if isConsonant(input) {
		return "Konsonan"
	} else if isVowel(input) {
		return "Vokal"
	}
	return "Unknown"
}

func main() {
	var input string
	fmt.Scanf("%s", &input)
	fmt.Println(checkVowelOrConsonant(input))
}
