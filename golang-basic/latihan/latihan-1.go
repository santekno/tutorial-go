package main

import (
	"bufio"
	"fmt"
	"os"
)

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Print("Welcome to Palindrome check, please input text to check is it a palindrome or not. to exit the program please type 'end' in the input field \n")

	// for is to make a loop so the program will always running until it stop by the user
	for {
		// scan input from terminal
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Enter a string: ")
	
		scanner.Scan()
	
		input := scanner.Text()

		// type end to end the program
		if input == "end" {
			break
		}

		// empty string validation
		inputNotEmpty := true

		if input == "" {
			fmt.Println("Input cannot be an empty string")
			inputNotEmpty = false
		}
	
		if inputNotEmpty {
			palindromeValidation := isPalindrome(input)
	
			if palindromeValidation {
				fmt.Println("Palindrome")
			} else {
				fmt.Println("Not a Palindrome")
			}
		}
	}
}
