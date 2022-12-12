package main

import "fmt"

func main() {

	fmt.Print("Enter a input: ")
	var input string
	fmt.Scanf("%s", &input)

	fmt.Println(palindrome(input))
}

func palindrome(word string) bool {
	res := []byte{}

	for i := len(word) - 1; i >= 0; i-- {
		res = append(res, word[i])
	}

	return word == string(res)
}
