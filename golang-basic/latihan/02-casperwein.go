package main

import "fmt"

func main() {
	fmt.Println(isKonsonan("a"))
	fmt.Println(isKonsonan("h"))
}

func isKonsonan(h string) string {
	if h == "a" || h == "i" || h == "u" || h == "e" || h == "o" {
		return "vokal"
	} else {
		return "konsonan"
	}
}
