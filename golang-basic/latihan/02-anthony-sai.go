package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)

	if input == "a" || input == "e" || input == "i" || input == "o" || input == "u" {
		fmt.Println("Vocal.")
	} else {
		fmt.Println("Consonant.")
	}
}
