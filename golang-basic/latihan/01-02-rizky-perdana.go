package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	menu()
}

func menu() {
	var menuInput string = ""

	for menuInput != "99" {
		fmt.Println("TASK EDSPERT GOLANG DAY 1")
		fmt.Println("Choose what service you want to use: ")
		fmt.Println("1. Check Polindrome")
		fmt.Println("2. Check Vokal or Konsonan")
		fmt.Println("99. Exit")
		fmt.Print("Input menu :\t")
		fmt.Scan(&menuInput)

		if validateInputMenu(menuInput) {
			switch menuInput {
			case "1":
				menuPolindrome()
			case "2":
				menuConsonant()
			case "99":
				fmt.Println("Thank you")
			default:
				fmt.Println("Please choose the correct menu")
			}
		}

		fmt.Println()
		fmt.Println()

	}

}

func menuPolindrome() {
	var textInput string

	fmt.Println()
	fmt.Println("Menu polindrome")
	fmt.Print("Input text :\t")
	fmt.Scan(&textInput)

	if !validateInputText(textInput) {
		result := checkPolindrome(textInput)
		fmt.Println(result)
		fmt.Println()
		fmt.Println()
	}

}

func menuConsonant() {
	var textInput string

	fmt.Println()
	fmt.Println("Menu consonant")
	fmt.Print("Input text :\t")
	fmt.Scan(&textInput)

	if !validateInputText(textInput) {
		if len(textInput) == 1 {
			result := checkConsonant(textInput)
			fmt.Println(result)
			fmt.Println()
			fmt.Println()
		} else {
			fmt.Println("Text input must only one character")
			fmt.Println()
			fmt.Println()
		}
	}
}

func checkPolindrome(textInput string) string {
	var textInputUpperCase string = strings.ToUpper(textInput)
	var halfLengthText int = len(textInputUpperCase) / 2

	for i := 0; i < halfLengthText; i++ {
		if textInputUpperCase[i] != textInputUpperCase[len(textInputUpperCase)-i-1] {
			//fmt.Println("Text input is not polindrome")
			//return
			return "Text input is not polindrome"
		}
	}

	//fmt.Println("Text input is polindrome")
	return "Text input is polindrome"
}

func checkConsonant(textInput string) string {
	var alphabetConsonant string = "AEIOU"

	if strings.Contains(alphabetConsonant, strings.ToUpper(textInput)) {
		//fmt.Println("Text input is vocal")
		//return
		return "Text input is vocal"
	} else {
		//fmt.Println("Text input is consonant")
		return "Text input is consonant"
	}
}

func validateInputText(textInput string) bool {
	var validateInputText bool = false

	if inputIsBlank(textInput) {
		fmt.Println("Text input must not blank")
		fmt.Println()
		fmt.Println()
		validateInputText = true
	} else if !inputIsAlphabet(textInput) {
		fmt.Println("Text input must only contains alphabet")
		fmt.Println()
		fmt.Println()
		validateInputText = true
	}

	return validateInputText
}

func validateInputMenu(textInput string) bool {
	var validateInputText bool = true

	if inputIsBlank(textInput) {
		fmt.Println("Text input must not blank")
		fmt.Println()
		fmt.Println()
		validateInputText = false
	} else if !inputIsNumeric(textInput) {
		fmt.Println("Text input must only contains numeric")
		fmt.Println()
		fmt.Println()
		validateInputText = false
	}

	return validateInputText
}

func inputIsBlank(textInput string) bool {
	return strings.TrimSpace(textInput) == ""
}

func inputIsAlphabet(textInput string) bool {
	return regexp.MustCompile(`^[A-Za-z]+$`).MatchString(textInput)
}

func inputIsNumeric(textInput string) bool {
	return regexp.MustCompile(`\d+`).MatchString(textInput)
}
