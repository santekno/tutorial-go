package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Please input your words: ")
	fmt.Scan(&input)
	consonantOrVocal(input)
}

func consonantOrVocal(input string) {
	v := "aiueo"
	a := strings.ToLower(input)
	if len(a) > 1 || len(a) < 1 {
		fmt.Println("Please input just 1 letter!")
		return
	}

	if strings.Contains(v, a) {
		fmt.Println("Vokal")
		return
	}

	fmt.Println("Konsonan")
}
