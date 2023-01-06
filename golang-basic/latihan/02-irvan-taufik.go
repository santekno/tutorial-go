package main

import (
	"fmt"
	"strings"
)

func consonantOrVocal(letter string) string {
	letterTolower := strings.ToLower(letter)
	vocal := [...]string{"a", "i", "u", "e", "o"}

	for i := 0; i < len(vocal); i++ {
		if letterTolower == vocal[i] {
			return "vocal"
		}
	}
	return "consonant"

}

func main() {
	fmt.Println(consonantOrVocal("a"))
	fmt.Println(consonantOrVocal("k"))
	fmt.Println(consonantOrVocal("U"))

}
