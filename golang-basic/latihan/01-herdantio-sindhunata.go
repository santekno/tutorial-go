package main

import "fmt"

func isPalindrome(input string) bool {
	length := len(input)
	limit := length / 2

	for i := 0; i < limit; i++ {
		if input[i] != input[length-i-1] {
			return false
		}
	}

	return true
}
func main() {
	var input string
	var output bool

	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		return
	}

	output = isPalindrome(input)
	println(output)
}
