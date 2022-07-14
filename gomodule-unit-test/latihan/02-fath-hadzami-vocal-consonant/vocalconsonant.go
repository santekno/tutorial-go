package vocalconsonant

func IsLetter(char string) bool {
	for _, r := range char {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

func CheckChar(char string) string {
	if len(char) != 1 {
		return "Hanya boleh 1 karakter"
	}

	if !IsLetter(char) {
		return "Hanya boleh huruf"
	}

	switch char {
	case "a", "i", "u", "e", "o", "A", "I", "U", "E", "O":
		return "Vokal"
	}

	return "Konsonan"
}
