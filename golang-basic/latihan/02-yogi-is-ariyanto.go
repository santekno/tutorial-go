package main

import "fmt"

func main() {
	var char string
	fmt.Scanf("%s", &char)

	isCharacter(char)
}

func isCharacter(character string) {
	switch character {
	case "a", "i", "u", "e", "o":
		fmt.Printf("%s ini adalah vokal\n", character)
	case "b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "x", "y", "z":
		fmt.Printf("%s ini adalah konsonan\n", character)
	default:
		fmt.Printf("%s ini bukan vokal dan bukan konsonan\n", character)
	}
}
