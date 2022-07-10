package main

import "fmt"

func isPolindrome(words string) bool {
	kata := ""
	for _, word := range words {
		kata = string(word) + kata
	}
	if words == kata {
		return true
	} else {
		return false
	}
}

func main() {
	isPlindrome := isPolindrome("asdadasdasdaxczc")
	fmt.Println(isPlindrome)
}
