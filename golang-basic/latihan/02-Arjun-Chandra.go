package main

import (
	"fmt"
	"strings"
)

func checkTypeWord(word string) {
	if len(word) == 1 {
		vocal := "aiueo"
		consonant := "bcdefghijklmnopqrstuvwxyz"
		convertedWord := strings.ToLower(word)
		if strings.Contains(vocal, convertedWord) {
			fmt.Println("Ini Adalah Huruf Vokal")
		} else if strings.Contains(consonant, word) {
			fmt.Println("Ini Adalah Huruf Konsonan ")
		} else {
			fmt.Println("Ini Bukan Huruf Vocal Ataupun Huruf Konsonan")
		}
	} else {
		fmt.Println("anda hanya boleh memasukan 1 huruf")
	}
}

func main() {
	var input string
	fmt.Print("Masukan Huruf : ")
	fmt.Scanf("%s", &input)
	checkTypeWord(input)
}
