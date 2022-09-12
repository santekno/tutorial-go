package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("masukan huruf alfabet ")
	var input_string string
	fmt.Scanln(&input_string)
	fmt.Println(Check_string(input_string))

}
func Check_string(input_string string) string {
	make_lower_string := strings.ToLower(input_string)
	if len(input_string) == 0 {
		return "alfabet tidak boleh kosong"

	} else {
		if len(input_string) > 1 {
			return "alfabet harus 1 huruf"
		} else {
			if make_lower_string == "a" || make_lower_string == "i" || make_lower_string == "u" || make_lower_string == "e" || make_lower_string == "o" {
				return make_lower_string + " adalah vokal"
			} else {
				return make_lower_string + " adalah konsonan"
			}
		}
	}

}
