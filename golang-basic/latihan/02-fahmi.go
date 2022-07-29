package main

import (
	"fmt"
	"regexp"
)
	
func main()  {
	var input string
	fmt.Scanf("%s", &input)
	

	reconsonan := regexp.MustCompile(`[^aeiouAEIOU0-9\W]+`)
	revocal := regexp.MustCompile(`[aeiouAEIOU]+`)
	simbol := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	 
	 if reconsonan.MatchString(input) {
		fmt.Println("Consonan")
	} else if revocal.MatchString(input) {
		fmt.Println("vocal")
	} else if simbol.MatchString(input) {
		fmt.Println("simbol")
	} else {
		fmt.Println("Angka")
	}
	 // true
}