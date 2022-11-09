package main

import "fmt"

func main() {
	isVocal('a')
	isVocal('k')
}
func isVocal(character rune) {
	if character == 'a' || character == 'i' || character == 'u' || character == 'e' || character == 'o' {
		fmt.Printf("%c adalah vocal\n", character)
	} else {
		fmt.Printf("%c adalah konsonan\n", character)
	}
}
