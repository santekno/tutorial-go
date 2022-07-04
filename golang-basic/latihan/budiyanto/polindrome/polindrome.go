package main

import (
	"fmt"
	"strings"
)

func checkPolindrome(input string) bool {
	pecah := strings.Split(input, "")
	length := len(pecah)
	var leftIndex, rightIndex int
	if length%2 == 0 {
		leftIndex = length / 2
		rightIndex = leftIndex + 1
	} else {
		leftIndex = (length - 1) / 2
		rightIndex = leftIndex + 2
	}
	leftIndex -= 1
	rightIndex -= 1
	for i := 0; i <= leftIndex; i++ {
		if pecah[leftIndex-i] != pecah[rightIndex+i] {
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
