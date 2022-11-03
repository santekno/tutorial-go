package main

import (
	"fmt"
	"strings"
)

const (
	vokal    = "aiueo"
	konsonan = "bcdfghjklmnpqrstvwxyz"
)

func main() {
	var alphabet string
	//! Input alphabet via terminal
	fmt.Scanf("%s", &alphabet)
	fmt.Println(isKonsonanOrVokal(alphabet))
}

func isKonsonanOrVokal(substr string) string {
	//! conditions according to what has been entered by the user
	if strings.Contains(vokal, strings.ToLower(substr)) {
		return "Vokal"
	} else if strings.Contains(konsonan, strings.ToLower(substr)) {
		return "Konsonan"
	}

	return "input invalid"
}
