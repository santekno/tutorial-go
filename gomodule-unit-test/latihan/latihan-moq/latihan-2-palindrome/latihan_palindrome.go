package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)
	fmt.Println(palindrome(input))
}

func palindrome(inp string) bool {
	strlength := len(inp)
	var out string
	out = ""
	for i := strlength - 1; i >= 0; i-- {
		out = out + string(inp[i])
	}
	// fmt.Println(out)
	return inp == out
}
