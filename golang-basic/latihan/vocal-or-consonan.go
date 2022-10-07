package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Masukkan huruf: ")
	fmt.Scan(&input)

	fmt.Println(isVocalOrConsonant(input))
}

func isVocalOrConsonant(str string) string {
	if len(str) > 1 {
		return "Input harus 1 huruf"
	}

	_, err := strconv.Atoi(str)
	if err == nil {
		return "Input harus huruf"
	}

	if str == "a" || str == "i" || str == "u" || str == "e" || str == "o" {
		return "Vocal"
	}
	return "Konsonan"
}
