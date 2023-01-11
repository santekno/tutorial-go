package main

import "fmt"

func main() {

	var str string

	fmt.Scan(&str)

	result := isVocalKonsonan(str)

	fmt.Println(result)
}

func isVocalKonsonan(str string) string {
	if str == "a" || str == "e" || str == "i" || str == "o" || str == "u" || str == "A" || str == "E" || str == "I" || str == "O" || str == "U" {
        return "Vokal"
    } else {
        return "Konsonan"
    }
}