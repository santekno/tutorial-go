package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string

Check:
	fmt.Print("Input string vocal/konsonan : ")
	fmt.Scan(&word)
	if len(word) != 1 {
		fmt.Println("!! Inputan tidak valid harus satu hurup !!")
		goto Check
	}

	fmt.Println(isVocal(word))
}

func isVocal(character string) string {
	vocals := []string{"a", "i", "u", "e", "o"}

	for _, r := range vocals {
		if strings.ToLower(r) == strings.ToLower(character) {
			return "Vocal"
		}
	}

	return "Konsonan"
}
