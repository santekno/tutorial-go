package main

import (
	"fmt"

	"4d63.com/strrev"
)

func main() {
	fmt.Print("masukan string ")
	var polidrome string
	fmt.Scanln(&polidrome)

	hasil := strrev.Reverse(polidrome)

	if polidrome == hasil {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
