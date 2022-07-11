package main

import "fmt"

func Palindrom(word string, i int) string{
	lengthWord := len(word)
	fmt.Print(lengthWord/2)
	if i < lengthWord/2 {
		if string(word[i]) != string(word[lengthWord - i - 1]){
			return "false"
		} else {
			Palindrom(word, i+1)
		}
	}
	return "true" 
}

func main() {
	var word string
	fmt.Scan(&word)
	fmt.Print(" ")
	fmt.Println(Palindrom(word, 0))
}