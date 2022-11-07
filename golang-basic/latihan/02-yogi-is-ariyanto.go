package main

import "fmt"

func main() {
	isCharacter("a")
	isCharacter("b")
	isCharacter("c")
	isCharacter("d")
	isCharacter("e")
	isCharacter("i")
	isCharacter("z")
}

func isCharacter(character string) {
	switch character {
	case "a", "i", "u", "e", "o":
		fmt.Printf("%s ini adalah vokal\n", character)
	default:
		fmt.Printf("%s ini adalah konsonan\n", character)
	}
}
