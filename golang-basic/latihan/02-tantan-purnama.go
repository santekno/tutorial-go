package main

import (
	"fmt"
	"regexp"
)

func Check(s string) {
	r, _ := regexp.Compile("[auieoAUIEO]")
	rs, _ := regexp.Compile("[a-zA-Z]")
	for _, x := range s {
		if rs.MatchString(string(x)) {
			if r.MatchString(string(x)) {
				fmt.Printf("%s Vokal\n", string(x))
			} else {
				fmt.Printf("%s Konsonan\n", string(x))
			}
		} else {
			fmt.Printf("%s not char\n", string(x))
		}
	}
}

func main() {
	var input string

	fmt.Print("Enter a string: ")
	fmt.Scanf("%s", &input)
	Check(input)
}
