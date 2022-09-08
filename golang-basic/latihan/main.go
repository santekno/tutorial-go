package main

import "fmt"

func main() {
	var choice int
	fmt.Println("Choice function")
	fmt.Println("1. Check palindrome")
	fmt.Println("2. Check consonant letter")
	fmt.Scanln(&choice)
	choice_map := map[int]func(){
		1: main_01,
		2: main_02,
	}

	fun, ok := choice_map[choice]
	if ok {
		fun()
	}
}
