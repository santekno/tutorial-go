package main

import (
	"fmt"
	"strings"
)

//konsonan

func vokalDetection(x string) string {
	var hurufVokal = [5]string{"a", "i", "u", "e", "o"}
	var huruf = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var isHuruf bool
	var jawaban string
	jawaban = "Konsonan"
	x = strings.ToLower(x)
	if len(x) == 1 {
		for i := 0; i < len(huruf); i++ {
			if huruf[i] == x {
				isHuruf = true
			}
		}
		if isHuruf {
			for i := 0; i < len(hurufVokal); i++ {
				if hurufVokal[i] == x {
					jawaban = "Vokal"
				}
			}
		} else {
			jawaban = "Inputan Bukan Huruf, Tolng Input Huruf"
		}
	} else {
		jawaban = "Tolong input 1 karakter saja"
	}

	return jawaban
}

func main() {
	var inputHuruf string
	fmt.Scan(&inputHuruf)
	fmt.Println(vokalDetection(inputHuruf))
}
