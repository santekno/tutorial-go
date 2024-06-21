package main

import "fmt"

func main() {
	pol1 := "madam"
	pol2 := "kuda"

	fmt.Println(Polindrom(pol1)) // true
	fmt.Println(Polindrom(pol2)) // false
}

func Polindrom(pol string) bool {
	len := len(pol)
	for i := 0; i < len/2; i++ {
		if pol[i] != pol[len-i-1] {
			return false
		}
	}
	return true
}
