package latihantest

import (
	"strings"
)

func vocalcheck(dicek string) string {
	validasi := strings.ToLower(dicek)
	switch validasi {
	case "a", "i", "u", "e", "o":
		return ("huruf Vokal")

	default:
		return ("huruf Konsonan")
	}

}
