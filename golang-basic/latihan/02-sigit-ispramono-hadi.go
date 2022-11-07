package main

import "fmt"

//Soal 2 : Cek Huruf Vokal atau Konsonan
func cekVokalKonsonan(kata string) string {
	vokal := []string{"a", "i", "u", "e", "o"}
	konsonan := []string{"b", "c", "d", "f", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}
	numerik := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	if len(kata) > 1 {
		return "lebih dari satu"
	}
	for i := 0; i < len(kata); i++ {
		for j := range vokal {
			if string(kata[i]) == vokal[j] {
				return "huruf vokal"
			}
		}
		for k := range konsonan {
			if string(kata[i]) == konsonan[k] {
				return "huruf konsonan"
			}
		}
		for l := range numerik {
			if string(kata[i]) == numerik[l] {
				return "angka"
			}
		}
	}
	return "bukan huruf atau angka"
}

func main() {
	var masukan2 string
	fmt.Println("Masukkan huruf : ")
	fmt.Scanln(&masukan2)
	hasil2 := cekVokalKonsonan(masukan2)
	if hasil2 == "lebih dari satu" {
		fmt.Println("Masukkan satu huruf saja")
	} else if hasil2 == "angka" {
		fmt.Println("Masukkan huruf, bukan angka")
	} else if hasil2 == "huruf vokal" || hasil2 == "huruf konsonan" {
		fmt.Println("Huruf yang dimasukkan adalah", hasil2)
	} else {
		fmt.Println("Masukkan huruf vokal atau konsonan")
	}
}
