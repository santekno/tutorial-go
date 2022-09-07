package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func IsPalindrome(value string) bool {
	chars := []rune(value)
	var result []rune
	for i := len(chars) - 1; i >= 0; i-- {
		result = append(result, chars[i])
	}
	if string(result) == value {
		return true
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukan 1 suku kata: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	fmt.Println(IsPalindrome(input))
}
