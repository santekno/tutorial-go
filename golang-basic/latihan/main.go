package main

import "fmt"

func main() {
	var input, input2 string

	fmt.Println("Cek Palindrom")
check1:
	fmt.Print("Masukan sebuah kata : ")
	fmt.Scanf("%s\n", &input)
	if len(input) > 1 {
		if isPalindrom(input) {
			fmt.Println(input, "adalah polindrom")
		} else {
			fmt.Println(input, "bukanlah polindrom")
		}
	} else {
		fmt.Println("Masukanlah string yang benar, minimal 2 karakter")
		goto check1
	}
	fmt.Println("")
	fmt.Println("Cek vokal atau konsonan")
check2:
	fmt.Print("Masukan sebuah huruf (A-Z atau a-z) : ")
	fmt.Scanf("%s\n", &input2)
	if len(input2) == 1 {
		hasil := cekKonsonanVokal(input2)
		if hasil == 0 {
			fmt.Println("Inputan tidak sesuai")
			goto check2
		} else if hasil == 1 {
			fmt.Println(input2, "adalah huruf vokal")
		} else {
			fmt.Println(input2, "adalah huruf konsonan")
		}
	} else if len(input2) == 0 {
		fmt.Println("Anda belum memasukan huruf")
		goto check2
	} else {
		fmt.Println("Jumlah huruf lebih dari satu")
		goto check2
	}
}

func isPalindrom(kata string) bool {
	var hasil bool = true // Default adalah polindrom
	var z int = len(kata) - 1
	for a := 0; a <= (len(kata)/2)-1; a++ { // Bandingkan depan dan belakang, dan seterusnya hingga tengah
		if kata[a] != kata[z] { // Jika ditemukan beda
			hasil = false // Maka bukan polindrom
			break
		}
		z--
	}
	return hasil
}

func cekKonsonanVokal(karakter string) int {
	var hasil int = 0 // Default adalah bukan huruf
	// Apakah huruf A-Z atau a-z?
	if (int(karakter[0]) >= 65) && (int(karakter[0]) <= 90) || (int(karakter[0]) >= 97) && (int(karakter[0]) <= 122) {
		switch string(karakter[0]) {
		case "A", "I", "U", "E", "O", "a", "i", "u", "e", "o":
			hasil = 1 // Huruf vokal
		default:
			hasil = 2 // Huruf konsonan
		}
	}
	return hasil
}
