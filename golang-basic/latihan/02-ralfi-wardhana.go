package main

import (
	"fmt"
	"strings"
)

func konsonanOrVocal(huruf string) string {

	convertLowerCase := strings.ToLower(huruf)

	var result string

	if len(convertLowerCase) > 1 {
		result = "Hanya bisa masukan 1 huruf bukan kata atau kalimat"
		return result
	}

	switch convertLowerCase {
	case "a":
		result = "Vocal"
	case "i":
		result = "Vocal"
	case "u":
		result = "Vocal"
	case "e":
		result = "Vocal"
	case "o":
		result = "Vocal"
	default:
		result = "Konsonan"
	}

	return result
}

func main() {

	result := konsonanOrVocal("e")
	fmt.Println(result)
}
