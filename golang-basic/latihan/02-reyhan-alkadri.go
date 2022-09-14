package main

import "fmt"

func isVokal(character rune) {
	if character == 'a' || character == 'i' || character == 'u' || character == 'e' || character == 'o' {
		fmt.Printf(" %c adalah huruf vokal\n", character)
	} else {
		fmt.Printf(" %c adalah huruf konsonan\n", character)
	}

}

func main() {
	isVokal('o') // vokal
	isVokal('k') // konsonan
}
