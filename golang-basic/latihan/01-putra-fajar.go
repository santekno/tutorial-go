package main

import (
	"strings"
)

// func main() {
// 	var input string
// 	fmt.Println("Please input words: ")
// 	fmt.Scan(&input)
// 	res := isPalindrome(input)
// 	if res {
// 		fmt.Println("true")
// 	} else {
// 		fmt.Println("false")
// 	}
// }

func isPalindrome(input string) bool {
	a := strings.ToLower(input)
	b := []byte(a)
	for i := 0; i < len(b)/2; i++ {
		if a[i] != a[len(b)-1-i] {
			return false
		}
	}

	return true
}
