package main

import (
	"fmt"
)

func main() {
	// Soal 2
	fmt.Print("Enter a letter : ")
	var input string
	fmt.Scanf("%s", &input)

	switch input {
	case "a", "i", "u", "e", "o":
		fmt.Println("Vokal")
	default:
		fmt.Println("Konsonan")
	}
}
