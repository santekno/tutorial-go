package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)

	var inputCopy string

	for i := len(input) - 1; i >= 0; i-- {
		inputCopy += string(input[i])
	}

	if input == inputCopy {
		fmt.Println("Palindrome.")
	} else {
		fmt.Println("Not Palindrome.")
	}
}
