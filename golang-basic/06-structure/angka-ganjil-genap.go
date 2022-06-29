package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " Genap")
		} else {
			fmt.Println(i, " Ganjil")
		}
	}
}
