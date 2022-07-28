package main

import (
	"fmt"
)

func checkPolindrome(input string) bool {
	var (
		length     int = len(input)
		rightIndex int = length / 2
		leftIndex  int = rightIndex - 1
	)
	if rightIndex+rightIndex != length {
		rightIndex = leftIndex + 2
	}
	for i := 0; i <= leftIndex; i++ {
		if input[leftIndex-i] != input[rightIndex+i] {
			return false
		}
	}
	return true
}

func main() {
	var input string
	fmt.Scanf("%s", &input)
	fmt.Println(checkPolindrome(input))
}
