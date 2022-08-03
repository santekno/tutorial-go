package main

import (
	"fmt"
	"strings"
)

func main() {
	var alphabet string
	fmt.Println("Masukkan huruf: ")
	fmt.Scanln(&alphabet)

	if len(alphabet) > 1 {
		fmt.Println("Masukkan hanya 1 huruf.")
		return
	}

	var strLower = strings.ToLower(alphabet)

	if strLower == "a" || strLower == "i" || strLower == "u" || strLower == "e" || strLower == "o" {
		fmt.Printf("%s adalah Huruf Vokal.", alphabet)
	} else {
		fmt.Printf("%s adalah Huruf Konsonan.", alphabet)
	}
}
