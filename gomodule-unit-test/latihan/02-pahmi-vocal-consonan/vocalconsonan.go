package main

import (
	"fmt"
	"regexp"
)

func Vocal(input string) string {
	reconsonan := regexp.MustCompile(`[^aeiouAEIOU0-9\W]+`)
	revocal := regexp.MustCompile(`[aeiouAEIOU]+`)
	simbol := regexp.MustCompile(`[^a-zA-Z0-9]+`)

	if reconsonan.MatchString(input) {
		return "Consonan"
	} else if revocal.MatchString(input) {
		return "vocal"
	} else if simbol.MatchString(input) {
		return "simbol"
	} else {
		return "Angka"
	}

}

func main() {
	var input string
	fmt.Print("Enter vocal/consonan : ")
	fmt.Scanf("%s", &input)
	fmt.Println(Vocal(input))

	// true
}
