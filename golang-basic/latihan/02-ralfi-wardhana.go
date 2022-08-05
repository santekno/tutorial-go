package main

import (
	"fmt"
	"strings"
)

func konsonanOrVocal(huruf string) string {

	convertLowerCase := strings.ToLower(huruf)
	convertToArray := strings.Split(convertLowerCase, "")
	split := convertToArray[:]

	var result string

	if len(split) > 1 {
		result = "Hanya bisa masukan 1 huruf bukan kata atau kalimat"
		return result
	}
	for _, value := range split {
		switch value {
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
	}
	return result
}

func main() {

	result := konsonanOrVocal("a")
	fmt.Println(result)
}
