package main

import (
	"fmt"
)

// func isVokal(character rune) {
// 	if character == 'a' || character == 'e' || character == 'i' || character == 'o' || character == 'u' {
// 		fmt.Printf("%c adalah vokal\n", character)
// 	} else {
// 		fmt.Printf("%c adalah konsonan\n", character)
// 	}

// }

func isVokal(character rune) {
	switch character {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Printf(" %c adalah vokal\n", character)
	default:
		fmt.Printf(" %c adalah konsonan\n", character)
	}
}

func main() {
	var karakter rune
	fmt.Scanf("%c", &karakter)
	isVokal(karakter)
}
