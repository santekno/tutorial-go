package helper

import (
	"fmt"
	"strings"
)

func IsVokal(input string) string {
	// masukan semua huruf vokal ke variable, dan lakukan concat huruf vokal yang sudah di ubah menjadi Upper Case
	vokal := "aiueo"
	vokal += strings.ToUpper(vokal)

	// masukan semua huruf konsonan ke variable, dan lakukan concat huruf konsonan yang sudah di ubah menjadi Upper Case
	konsonan := "bcdfghjklmnpqrstvwxyz"
	konsonan += strings.ToUpper(konsonan)

	//	user hanya boleh input 1 huruf, jika kurang atau lebih return dan berikan pesan
	if len(input) < 1 || len(input) > 1 {
		return "Inputan tidak boleh kurang dari atau lebih dari 1"
	}

	// jika yang di input user valid, cek apakah inputan vokal atau konsonan
	// jika user menginput huruf atau simbol maka return dan beritahu jika yang di input bukan vokal atau konsonan
	if strings.Contains(vokal, input) {
		return fmt.Sprintf("%s adalah vokal", input)
	} else if strings.Contains(konsonan, input) {
		return fmt.Sprintf("%s adalah konsonan", input)
	} else {
		return fmt.Sprintf("%s bukan vokal ataupun konsonan", input)
	}
}
