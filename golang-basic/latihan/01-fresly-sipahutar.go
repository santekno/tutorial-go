package main

import "fmt"

func main() {
	var textPolindrome string
	fmt.Println("Masukkan kata: ")
	fmt.Scanln(&textPolindrome)

	reverseText := reverseStr(textPolindrome)

	if textPolindrome == reverseText {
		fmt.Println(textPolindrome == reverseText)
	} else {
		fmt.Println(textPolindrome == reverseText)
	}
}

func reverseStr(str string) string {
	revStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		revStr += string(str[i])
	}
	return revStr
}
