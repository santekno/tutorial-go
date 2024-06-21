package main

import (
	"fmt"
	"strings"
)

// Needed to use Split

func isPalindrom(str string) bool {
	arrayStr := []string{}
	strSplit := strings.Split(str, "")

	for i := len(strSplit) - 1; i >= 0; i-- {
		arrayStr = append(arrayStr, strSplit[i])
	}
	revesedStr := strings.Join(arrayStr, "")

	return revesedStr == str

}

func main() {

	fmt.Println(isPalindrom("kodok"))
	fmt.Println(isPalindrom("irvan"))
	fmt.Println(isPalindrom("ada"))
	fmt.Println(isPalindrom("madam"))
	fmt.Println(isPalindrom("katak"))

}
