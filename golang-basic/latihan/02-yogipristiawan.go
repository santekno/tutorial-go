package main

import (
	"errors"
	"fmt"
	"regexp"
)

func main() {
	var str string

	for true {
		// get user input
		getInput(&str)

		// validate input
		if err := validateInput(str); err != nil {
			fmt.Println(err.Error())
		} else {
			// check vowel
			if checkVowel(str) {
				fmt.Println("Vokal")
			} else {
				fmt.Println("Konsonan")
			}
			return
		}
	}
}

func getInput(str *string) {
	fmt.Println("\nCheck a string is vowel or consonant")
	fmt.Print("===============================", "\n\n")
	fmt.Print("Input some string: ")
	fmt.Scanln(str)
}

func validateInput(str string) (err error) {
	// validate length
	if len(str) > 1 {
		return errors.New("ERROR: input must be 1 character long")
	}

	// validate alphabet
	r, _ := regexp.Compile("^[a-zA-Z]$")
	if !r.MatchString(str) {
		return errors.New("ERROR: input must be an alphabet")
	}

	return nil
}

func checkVowel(str string) (isVowel bool) {
	const vowel string = "aiueoAIUEO"

	for _, val := range vowel {
		if string(val) == str {
			return true
		}
	}

	return
}
