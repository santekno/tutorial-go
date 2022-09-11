package main

import "fmt"

func main() {
	var choice int
	fmt.Println("Choice function")
	fmt.Println("1. Check palindrome")
	fmt.Println("2. Check consonant letter")
	fmt.Scanln(&choice)
	choiceMap := map[int]func(){
		1: main01,
		2: main02,
	}

	fun, ok := choiceMap[choice]
	if ok {
		fun()
	}
}
