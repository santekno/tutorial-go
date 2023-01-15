package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func isVowel(s string) bool {
	input := strings.ToLower(s)
	if(input == "a" || input == "i" || input == "u" || input == "e" || input == "o") {
		return true
	}
	return false
}

func isAlphabet(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z]*$`)

	return re.MatchString(s)
}

func main() {
	fmt.Print("Welcome to Vowel check, please input an alphabet to check is it a vowel or not. to exit the program please type 'end' in the input field \n")

	// for is to make a loop so the program will always running until it stop by the user
	for {
		// scan input from terminal
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Enter a single string: ")
	
		scanner.Scan()
	
		input := scanner.Text()

		// type end to end the program
		if input == "end" {
			break
		}

		// empty string validation & character validation
		inputValidationOk := true

		if input == "" || len(input) > 1 || !isAlphabet(input){
			fmt.Println("Input only can an alphabet and cannot be empty")
			inputValidationOk = false
		}


		if inputValidationOk {
			characterCheck := isVowel(input)

			if characterCheck {
				fmt.Println("Vowel")
			} else {
				fmt.Println("Not a Vowel")
			}
		}
	}
}