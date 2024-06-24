/* Soal 2
	Buat program untuk menentukan Huruf yang diinput Vokal atau Konsonan
 >>> input
	$ go main .go
	h
 >>> output
    Konsonan
*/

package main

import "fmt"

func isVocal(vocal string) string {

	if vocal == "a" || vocal == "e" || vocal == "i" || vocal == "o" || vocal == "u" {
		vocal = "Vokal"
	} else {
		vocal = "Konsonan"
	}

	return vocal
}

func main() {
	var vocal string
	fmt.Print("Masukkan huruf : ")
	fmt.Scan(&vocal)

	result := isVocal(vocal)

	if result == "Vokal" {
		fmt.Println("Vokal")
	} else {
		fmt.Println("Konsonan")
	}

}
