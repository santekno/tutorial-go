package pkg

func IsVocal(char string) string {
	if (char == "a") || (char == "i") || (char == "u") || (char == "e") || (char == "o") {
		return "Vocal"
	} else {
		return "Consonant"
	}
}
