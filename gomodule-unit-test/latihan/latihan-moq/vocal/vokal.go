package vokal

import (
	"fmt"
	"strings"
)

func main() {
	var inputChar string
	fmt.Println("Input Character : ")
	fmt.Scanf("%s", &inputChar)
	isValidInput := validateInput(inputChar)
	if isValidInput {
		fmt.Println(IsVocal(inputChar))
	}
}

func validateInput(inputChar string) bool {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	if len(inputChar) != 1 {
		fmt.Println("The input character must be only 1 character..")
		return false
	} else if !strings.Contains(alpha, strings.ToLower(inputChar)) {
		fmt.Println("The input must be a string character..")
		return false
	}
	return true
}

func IsVocal(inputChar string) string {
	vocals := "aiueo"
	if !strings.Contains(vocals, inputChar) {
		return "Konsonan"
	}
	return "Vocal"
}
