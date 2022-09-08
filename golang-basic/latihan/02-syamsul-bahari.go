package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("masukan huruf alfabet ")
	var input_string string
	fmt.Scanln(&input_string)
	make_lower_string := strings.ToLower(input_string)
	if len(input_string) == 0 {
		fmt.Println("alfabet tidak boleh kosong")

	} else {
		if len(input_string) > 1 {
			fmt.Println("alfabet harus 1 huruf")
		} else {
			if make_lower_string == "a" || make_lower_string == "i" || make_lower_string == "u" || make_lower_string == "e" || make_lower_string == "o" {
				fmt.Printf("%s adalah vokal ", make_lower_string)
			} else {
				fmt.Printf("%s adalah konsonan ", make_lower_string)
			}
		}
	}

}
