package main

import "fmt"

// polindrone
func polindrone(x string) bool {
	var perbedaanHuruf byte
	perbedaanHuruf = 0
	var hasilAkhir bool
	for i := 0; i < len(x) && perbedaanHuruf == 0; i++ {
		perbedaanHuruf = x[i] - x[len(x)-i-1]
		if perbedaanHuruf == 0 {
			hasilAkhir = true
		} else {
			hasilAkhir = false
		}
	}
	return hasilAkhir
}

func main() {
	var inputKata string
	fmt.Scan(&inputKata)
	fmt.Println(polindrone(inputKata))
}
