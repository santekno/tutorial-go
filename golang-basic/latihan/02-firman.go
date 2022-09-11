package main

import (
	"fmt"
	"strings"
)

func consonantsOrVowels(str *string) string {
	*str = strings.ToLower(*str)
	if *str == "a" || *str == "i" || *str == "u" || *str == "e" || *str == "o" {
		return "Vokal"
	}
	return "Konsonan"
}

func main() {
	fmt.Print("Masukkan sebuah Huruf: ")
	var input string
	fmt.Scanln(&input)
	fmt.Println(consonantsOrVowels(&input))
}
