/*
Edspert Bootcamp - Golang Backend
Tugas 2 : Golang Basic
Peserta : Rizdin Ahmad Farizi
*/

package main

import "fmt"

func main() {
	words := ("")

	fmt.Scan(&words)

	switch words {
	case "a", "i", "u", "e", "o":
		fmt.Println("Huruf Vocal")

	case "A", "I", "U", "E", "O":
		fmt.Println("Huruf Vocal")

	default:
		fmt.Println("Ini Consonant")
	}
}
