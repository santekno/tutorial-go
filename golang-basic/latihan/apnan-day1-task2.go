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
			fmt.Println("This is vocal word")
		} else if strings.Contains(consonant, word) {
			fmt.Println("This is consonant word")
		} else {
			fmt.Println("This is not vocal & consonant word")
		}
	} else {
		fmt.Println("sorry, you can only input one character")
	}
}

func main() {
	var input string 
	fmt.Print("Please input your word: ")
	fmt.Scanf("%s", &input)
	checkTypeWord(input)
}