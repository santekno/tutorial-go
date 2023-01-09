package main

import "fmt"

func main() {

	var str string

	fmt.Scan(&str)

	result := isPalindrome(str)

	fmt.Println(result)

}

func isPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
