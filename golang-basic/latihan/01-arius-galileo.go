package main

import "fmt"

// polindrone

func main() {
	var inputKata string
	var hasilAkhir bool
	var perbedaanHuruf byte
	perbedaanHuruf = 0
	fmt.Scan(&inputKata)
	for i := 0; i < len(inputKata) && perbedaanHuruf == 0; i++ {
		perbedaanHuruf = inputKata[i] - inputKata[len(inputKata)-i-1]
		if perbedaanHuruf == 0 {
			hasilAkhir = true
		} else {
			hasilAkhir = false
		}
	}

	fmt.Println(hasilAkhir)
}
