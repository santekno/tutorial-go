package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main(){
	input := "Rotavator";

	isPolindrome := func(word string) bool {
		res := true
		z := 0
		for i := len(word) - 1; i >= 1; i--{
			if !strings.EqualFold(string(word[i]),(string(word[z]))) {
				return false
			}	
			z++		
		}
		return res
	}

	fmt.Println(isPolindrome(input))

	Vowel()
}

func Vowel(){
	fmt.Println("Enter the word):")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	input = strings.TrimSuffix(input, "\n")

	for _, r := range input{
		currentChar := string(r)
		if unicode.IsLetter(r) {
			if strings.EqualFold(currentChar, "A") || 
				strings.EqualFold(currentChar, "B") ||
				strings.EqualFold(currentChar, "C") ||
				strings.EqualFold(currentChar, "D") ||
				strings.EqualFold(currentChar, "E") {
				fmt.Println("Char: ", currentChar, " is Vowel")
			}else{
				fmt.Println("Char: ", currentChar, " is consonant")
			}
		}
	}
}