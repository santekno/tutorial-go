package main

import "fmt"

func Reverse(s string) (temp string) {
	for _, x := range s {
		temp = string(x) + temp
	}
	return
}

func main() {
	var input string

	fmt.Print("Enter a string: ")
	fmt.Scanf("%s", &input)

	RevStr := Reverse(input)
	fmt.Println((input == RevStr))
}
