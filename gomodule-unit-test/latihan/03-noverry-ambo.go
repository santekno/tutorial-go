package main

func isVocal(vocal string) string {

	if vocal == "a" || vocal == "e" || vocal == "i" || vocal == "o" || vocal == "u" {
		vocal = "Vokal"
	} else {
		vocal = "Konsonan"
	}
	return vocal
}
