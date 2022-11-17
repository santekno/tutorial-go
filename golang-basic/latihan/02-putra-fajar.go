package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Please input your words: ")
	fmt.Scan(&input)
	isVokal(input)
}

func isVokal(input string) bool {
	v := "aiueo"
	k := "qwrtypsdfghjklzxcvbnm"
	a := strings.ToLower(input)
	if len(a) > 1 || len(a) < 1 {
		fmt.Println("Please input just 1 letter!")
		return false
	}

	if strings.Contains(v, a) {
		fmt.Println("Vokal")
		return true
	}

	if strings.Contains(k, a) {
		fmt.Println("Konsonan")
		return false
	}

	fmt.Println("Input not letter..")
	return false
}
