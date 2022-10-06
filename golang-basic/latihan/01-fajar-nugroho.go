package main

import (
	"fmt"
)

func main() {
	var polindrom1, polindrom2 string
	polindrom1 = "ramukumar"
	polindrom2 = "fajar"

	res1 := isPolindrom(polindrom1)
	fmt.Println(res1)

	res2 := isPolindrom(polindrom2)
	fmt.Println(res2)

}

func isPolindrom(param string) bool {
	var output string
	var matched bool

	matched = false
	for i := len(param) - 1; i >= 0; i-- {
		output += string(param[i])
	}

	if output == param {
		matched = true
	}
	return matched
}
