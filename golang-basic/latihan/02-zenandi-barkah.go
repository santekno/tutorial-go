package main

import "fmt"

func main() {
	var inputChar string

	fmt.Print("Input your character: ", inputChar)
	fmt.Scanf("%s", &inputChar)

	result := isVocal(inputChar)
	fmt.Print(result)
}

func isVocal(char string) string {
	if (char == "a") || (char == "i") || (char == "u") || (char == "e") || (char == "o") {
		return "Vocal"
	} else {
		return "Consonant"
	}
}
