package main

import "fmt"

func Konsonan(str string) bool {
	konsonan := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}
	isKonsonanan := make(map[string]bool)
	for _, v := range konsonan {
		isKonsonanan[v] = true
	}
	return isKonsonanan[str]
}

func main() {
	var word string
	fmt.Println("Masukan Konsonan")
	fmt.Scan(&word)

	checkKonsonan := Konsonan(word)
	if checkKonsonan == true {
		fmt.Println("Kata Konsonan")
	} else {
		fmt.Println("Kata Bukan Konsonan")
	}

}
