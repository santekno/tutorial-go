package main

import (
	"fmt"
	"github.com/ApnanJuanda/tutorial-go/helper"
)

func main() {
	var input string 
	fmt.Print("Please input your word: ")
	fmt.Scanf("%s", &input)
	fmt.Println(helper.CheckTypeWord(input))
}