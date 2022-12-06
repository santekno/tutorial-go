package main

import "fmt"

func isVokal(character string) bool {
	if character == "a" || character == "e" || character == "i" || character == "o" || character == "u" {
		return true

	} else {

		return false
	}
}

func main() {

	var str string

	fmt.Print("Tulis Sebuah Huruf Disini : ")
	fmt.Scan(&str)

	result := isVokal(str)

	if result == true {

		fmt.Println("Vokal")

	} else {

		fmt.Println("Konsonan")

	}

}
