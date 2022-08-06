package main

import "fmt"

func isPolindrome(value string) bool {
	for i := 0; i < len(value)/2; i++ {
		if value[i] != value[len(value)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Print("Enter your word: ")
	var input string
	fmt.Scanf("%s", &input)

	output := isPolindrome(input)
	fmt.Println(output)
}
