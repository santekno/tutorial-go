package main

import (
	"fmt"
	"github.com/ApnanJuanda/tutorial-go/helper"
)

func main() {
	var input string
	fmt.Print("Please input your word: ")
	fmt.Scanf("%s", &input)
	if (helper.IsPolyndrome(input)) {
		fmt.Println("This is Polyndrome")
	} else {
		fmt.Println("This is not Polyndrome")
	}
}	