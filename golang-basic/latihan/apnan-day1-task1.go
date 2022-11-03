package main

import "fmt"

func isPolyndrome(word string) bool {
	if len(word) > 0 {
		result := true
		limit := len(word)/2
		for i := 0; i < limit; i++ {
			if (word[i] != word[len(word)-i-1]) {
				result = false
				break
			}
		}
		return result
	} else {
		return false
	}
	
}

func main() {
	var input string 
	fmt.Print("Please input your word: ")
	fmt.Scanf("%s", &input)
	if (isPolyndrome(input)) {
		fmt.Println("This is Polyndrome")
	} else {
		fmt.Println("This is not Polyndrome")
	}
}	