package latihan

import (
	"fmt"
	"regexp"
)

func Check(s string) string {
	if len(s) > 1 {
		return "too many character"
	}

	r, _ := regexp.Compile("[auieoAUIEO]")
	rs, _ := regexp.Compile("[a-zA-Z]")

	if rs.MatchString(string(s)) {
		if r.MatchString(string(s)) {
			return "vocal"
		} else {
			return "konsonan"
		}
	}

	return "not char"
}

func vocal() {
	var input string

	fmt.Print("Enter a string: ")
	fmt.Scanf("%s", &input)
	Check(input)
}
