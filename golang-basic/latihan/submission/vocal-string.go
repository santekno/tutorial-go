package submission

import (
	"strings"
)

const VOCAL string = "Vocal"
const CONSONANT string = "Konsonan"

func IsVocal(character string) string {
	vocals := []string{"a", "i", "u", "e", "o"}

	for _, r := range vocals {
		if strings.EqualFold(character, r) {
			return VOCAL
		}
	}

	return CONSONANT
}
