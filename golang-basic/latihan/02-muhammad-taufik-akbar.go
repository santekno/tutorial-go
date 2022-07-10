package main

import "fmt"

func isVokal(word string) bool {
	return word == "a" || word == "i" || word == "u" || word == "e" || word == "o" || word == "A" || word == "I" || word == "U" || word == "E" || word == "O"
}

func charChecker(word string) string {
	wordLength := len(word)
	if wordLength != 1 {
		return "INVALID"
	}
	if isVokal(word) {
		return "Vokal"
	}
	return "Konsonan"
}

func main() {
	var word string
	fmt.Print("Input: ")
	fmt.Scan(&word)
	fmt.Printf("Result: %s", charChecker(word))
}
