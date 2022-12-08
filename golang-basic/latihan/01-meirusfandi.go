package main

import "fmt"

func isPalindrom(pos int, str string) bool {
	if pos == len(str)-1 {
		return true
	} else if pos < len(str)-1 {
		if str[pos] == str[len(str)-pos-1] {
			return isPalindrom(pos+1, str)
		} else {
			return false
		}
	} else {
		return false
	}
}

func main() {
	var str string
	fmt.Print("Input some string : ")
	fmt.Scan(&str)
	fmt.Println(isPalindrom(0, str))
}
