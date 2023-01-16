package vocal

import (
	"strings"
)

func Vocal(letter string) string {
	letterTolower := strings.ToLower(letter)
	vocal := [...]string{"a", "i", "u", "e", "o"}

	for i := 0; i < len(vocal); i++ {
		if letterTolower == vocal[i] {
			return "vocal"
		}
	}
	return "consonant"

}
